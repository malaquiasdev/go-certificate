# ekoa-certificate-generator

## DynamoDB Schema

### METADATA
``` 
{
  "TableName": "report_enrollment",
  "KeySchema": [
    {
      "KeyType": "HASH",
      "AttributeName": "METADATA#0"
    },
    {
      "KeyType": "RANGE",
      "AttributeName": "METADATA#0"
    }
  ],
  "AttributeDefinitions": [
    {
      "AttributeName": "METADATA#0",
      "AttributeType": "S"
    },
    {
      "AttributeName": "totalCount",
      "AttributeType": "N"
    },
  ],
  "BillingMode": "PAY_PER_REQUEST"
}
```

### REPORT

```
{
  "TableName": "report_enrollment",
  "KeySchema": [
    {
      "KeyType": "HASH",
      "AttributeName": "REPORT#{ID}"
    },
    {
      "KeyType": "RANGE",
      "AttributeName": "MEMBER#{EMAIL}"
    }
  ],
  "AttributeDefinitions": [
    {
      "AttributeName": "MEMBER#{EMAIL}",
      "AttributeType": "S"
    },
    {
      "AttributeName": "REPORT#{ID}",
      "AttributeType": "N"
    }
  ],
  "BillingMode": "PAY_PER_REQUEST"
}
```