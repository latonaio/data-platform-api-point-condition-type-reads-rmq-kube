package dpfm_api_output_formatter

import (
	"data-platform-api-point-condition-type-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToPointConditionType(rows *sql.Rows) (*[]PointConditionType, error) {
	defer rows.Close()
	pointConditionType := make([]PointConditionType, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.PointConditionType{}

		err := rows.Scan(
			&pm.PointConditionType,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &pointConditionType, nil
		}

		data := pm
		pointConditionType = append(pointConditionType, PointConditionType{
			PointConditionType:		data.PointConditionType,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &pointConditionType, nil
}

func ConvertToText(rows *sql.Rows) (*[]Text, error) {
	defer rows.Close()
	text := make([]Text, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Text{}

		err := rows.Scan(
			&pm.PointConditionType,
			&pm.Language,
			&pm.PointConditionTypeName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &text, err
		}

		data := pm
		text = append(text, Text{
			PointConditionType:		data.PointConditionType,
			Language:          		data.Language,
			PointConditionTypeName:	data.PointConditionTypeName,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &text, nil
}
