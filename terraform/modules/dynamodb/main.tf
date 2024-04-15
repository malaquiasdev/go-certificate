resource "aws_dynamodb_table" "this" {
  count = var.aws_region == "sa-east-1" ? 1 : 0
  name  = var.name

  hash_key     = "PK"
  range_key    = "SK"
  billing_mode = "PAY_PER_REQUEST"

  point_in_time_recovery {
    enabled = var.point_in_time_recovery_enabled
  }

  attribute {
    name = "PK"
    type = "S"
  }

  attribute {
    name = "SK"
    type = "S"
  }

  attribute {
    name = "GS1PK"
    type = "S"
  }

  attribute {
    name = "GS1SK"
    type = "S"
  }

  attribute {
    name = "GS2PK"
    type = "S"
  }

  attribute {
    name = "GS2SK"
    type = "S"
  }

  global_secondary_index {
    name            = "GS1PK"
    hash_key        = "GS1PK"
    range_key       = "GS1SK"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "GS2PK"
    hash_key        = "GS2PK"
    range_key       = "GS2SK"
    projection_type = "ALL"
  }
}