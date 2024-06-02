package subscriber

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/maiquocthinh/go-comic/pkg/asyncjob"
	"github.com/maiquocthinh/go-comic/pkg/common"
	"github.com/maiquocthinh/go-comic/pkg/email"
	"github.com/maiquocthinh/go-comic/pkg/pubsub"
	"log"
)

type consumerJob struct {
	Title string
	Hld   func(ctx context.Context, msg string) error
}

type consumerEngine struct {
	ps       pubsub.PubSub
	db       *sqlx.DB
	emailSrv email.EmailService
}

func NewEngine(db *sqlx.DB, emailSrv email.EmailService, ps pubsub.PubSub) *consumerEngine {
	return &consumerEngine{db: db, emailSrv: emailSrv, ps: ps}
}

func (engine *consumerEngine) Start() {

	engine.startSubscribeTopic(
		common.TopicWriteHistoryView,
		false,
		WriteHistoryAfterViewChapter(engine.db),
	)
	engine.startSubscribeTopic(
		common.TopicIncreaseView,
		false,
		IncreaseViewAfterViewChapter(engine.db),
	)
	engine.startSubscribeTopic(
		common.TopicSendCodeResetPassword,
		false,
		SendCodeResetPasswordViaEmail(engine.emailSrv),
	)
}

func (engine *consumerEngine) startSubscribeTopic(topic string, isConcurrent bool, consumerJobs ...*consumerJob) {
	subListener := engine.ps.Subscribe(context.Background(), topic)

	go func() {
		defer subListener.UnSubscribe(context.Background())

		subListener.ListenAndReceive(func(msg interface{}) {
			jobHldArr := make([]asyncjob.Job, len(consumerJobs))

			for idx, _ := range consumerJobs {
				// wrap consumer job in async job
				var jobHl asyncjob.JobHandler = func(ctx context.Context) error {
					return consumerJobs[idx].Hld(context.Background(), msg.(string))
				}
				jobHldArr[idx] = asyncjob.NewJob(jobHl)
			}

			group := asyncjob.NewGroup(isConcurrent, jobHldArr...)
			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		})
	}()

}
