package uploadprovider

import (
	"bufio"
	"context"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"testing"
)

const (
	configPath = "../../config/config-local.yml"
)

type dropboxConfig struct {
	AppKey       string
	AppSecret    string
	RefreshToken string
}

func loadDropboxConfig() (*dropboxConfig, error) {
	v := viper.New()

	v.SetConfigFile(configPath)
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	config := struct{ Dropbox dropboxConfig }{}
	if err := v.Unmarshal(&config); err != nil {
		log.Printf("unable to decode into struct!")
		return nil, err
	}

	return &config.Dropbox, nil
}
func getBytesFromFile(filePath string) []byte {
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	// Get the file size
	stat, err := file.Stat()
	if err != nil {
		return nil
	}

	// Read the file into a byte slice
	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		return nil
	}

	return bs
}

func Test_UploadImage(t *testing.T) {
	config, err := loadDropboxConfig()
	if err != nil {
		t.Fatal("Load config fail.")
	}

	provider := &dropboxProvider{
		appKey:       config.AppKey,
		appSecret:    config.AppSecret,
		refreshToken: config.RefreshToken,
	}

	type args struct {
		ctx  context.Context
		data []byte
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Upload success",
			args{
				context.Background(),
				getBytesFromFile("./test_files/img.jpg"),
				"/test/upload/img.jpg",
			},
			false,
		},
		{
			"Upload success",
			args{
				context.Background(),
				getBytesFromFile("./test_files/text.txt"),
				"/test/upload/text.txt",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := provider.UploadImage(tt.args.ctx, tt.args.data, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("UploadImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_GetShareLink(t *testing.T) {
	config, err := loadDropboxConfig()
	if err != nil {
		t.Fatal("Load config fail.")
	}

	provider := &dropboxProvider{
		appKey:       config.AppKey,
		appSecret:    config.AppSecret,
		refreshToken: config.RefreshToken,
	}

	type args struct {
		ctx  context.Context
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"GetShareLink success",
			args{
				context.Background(),
				"/test/upload/img.jpg",
			},
			false,
		},
		{
			"GetShareLink success",
			args{
				context.Background(),
				"/test/upload/text.txt",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := provider.GetShareLink(tt.args.ctx, tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetShareLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" {
				t.Errorf("GetShareLink() got = %v, want %v", got, "string url")
			}
		})
	}
}
