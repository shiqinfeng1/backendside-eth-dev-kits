package command

import (
	"github.com/labstack/gommon/log"

	"github.com/shiqinfeng1/chunyuyisheng/service/common"
	"github.com/shiqinfeng1/chunyuyisheng/service/nsqs"
	"github.com/spf13/cobra"
)

type doctorResponseMock struct {
	common.ChunyuDoctorResponsePayload
}
type questionCloseMock struct {
	common.ChunyuQuestionClosePayload
}

var msgPusherCmd = &cobra.Command{
	Use:   "msgpusher",
	Short: "jpush tester: produce msgs.",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		common.Logger = log.New("msgpusher")
		if err := nsqs.InitConfig(); err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var doctorResponse = &doctorResponseMock{common.ChunyuDoctorResponsePayload{UserID: "123456", ProblemID: 654321}}
		var questionClose = &questionCloseMock{common.ChunyuQuestionClosePayload{UserID: "123456", ProblemID: 654321}}

		nsqs.PostTopic("topicChunyuDoctorResponse", doctorResponse)
		nsqs.PostTopic("topicChunyuQuestionClose", questionClose)
		return
	},
}
