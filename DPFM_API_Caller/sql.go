package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-point-condition-type-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-point-condition-type-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var pointConditionType *[]dpfm_api_output_formatter.PointConditionType
	var text *[]dpfm_api_output_formatter.Text
	for _, fn := range accepter {
		switch fn {
		case "PointConditionType":
			func() {
				pointConditionType = c.PointConditionType(mtx, input, output, errs, log)
			}()
		case "PointConditionTypes":
			func() {
				pointConditionType = c.PointConditionTypes(mtx, input, output, errs, log)
			}()
		case "Text":
			func() {
				text = c.Text(mtx, input, output, errs, log)
			}()
		case "Texts":
			func() {
				text = c.Texts(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		PointConditionType: pointConditionType,
		Text:      text,
	}

	return data
}

func (c *DPFMAPICaller) PointConditionType(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PointConditionType {
	where := fmt.Sprintf("WHERE PointConditionType = '%s'", input.PointConditionType.PointConditionType)

	if input.PointConditionType.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.PointConditionType.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_point_condition_type_point_condition_type_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, PointConditionType DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPointConditionType(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) PointConditionTypes(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.PointConditionType {
	where := fmt.Sprintf("WHERE PointConditionType = '%s'", input.PointConditionType.PointConditionType)

	if input.PointConditionType.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *input.PointConditionType.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_point_condition_type_point_condition_type_data
		` + where + ` ORDER BY IsMarkedForDeletion ASC, PointConditionType DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPointConditionType(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Text(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Text {
	var args []interface{}
	pointConditionType := input.PointConditionType.PointConditionType
	text := input.PointConditionType.Text

	cnt := 0
	for _, v := range text {
		args = append(args, pointConditionType, v.Language)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_point_condition_type_text_data
		WHERE (PointConditionType, Language) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToText(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Texts(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Text {
	var args []interface{}
	text := input.PointConditionType.Text

	cnt := 0
	for _, v := range text {
		args = append(args, v.Language)
		cnt++
	}

	repeat := strings.Repeat("(?),", cnt-1) + "(?)"
	rows, err := c.db.Query(
		`SELECT * 
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_point_condition_type_text_data
		WHERE Language IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	//
	data, err := dpfm_api_output_formatter.ConvertToText(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
