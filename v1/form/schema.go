package form

import (
	"cloud.google.com/go/bigquery"
)

// [
//     {
//         "name": "Email",
//         "type": "STRING",
//         "mode": "REQUIRED"
//     },
//     {
//         "name": "Name",
//         "type": "STRING",
//         "mode": "NULLABLE"
//     },
//     {
//         "name": "Subscribed",
//         "type": "BOOLEAN",
//         "mode": "NULLABLE"
//     },
//     {
//         "name": "Messages",
//         "type": "RECORD",
//         "mode": "REPEATED",
//         "fields": [
//             {
//                 "name": "Data",
//                 "type": "STRING",
//                 "mode": "NULLABLE"
//             },
//             {
//                 "name": "Timestamp",
//                 "type": "TIMESTAMP",
//                 "mode": "NULLABLE"
//             }
//         ]
//     }
// ]

var schema = bigquery.Schema{
	&bigquery.FieldSchema{Name: "Email", Type: bigquery.StringFieldType},
	&bigquery.FieldSchema{Name: "Name", Type: bigquery.StringFieldType},
	&bigquery.FieldSchema{Name: "Subscribed", Type: bigquery.BooleanFieldType},
	&bigquery.FieldSchema{Name: "Messages", Type: bigquery.RecordFieldType,
		Repeated: true,
		Schema: bigquery.Schema{
			&bigquery.FieldSchema{Name: "Data", Type: bigquery.StringFieldType},
			&bigquery.FieldSchema{Name: "Timestamp", Type: bigquery.TimestampFieldType},
		},
	},
}
