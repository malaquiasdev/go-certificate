resource "aws_sqs_queue" "main" {
  name          = var.name
  delay_seconds = var.delay_seconds
  visibility_timeout_seconds = var.visibility_timeout_seconds
  redrive_policy = jsonencode({
    deadLetterTargetArn = aws_sqs_queue.dlq.arn,
    maxReceiveCount     = var.max_redrive_count
  })
}

resource "aws_sqs_queue_policy" "main" {
  queue_url = aws_sqs_queue.main.id
  policy    = data.aws_iam_policy_document.main.json
}

resource "aws_sqs_queue" "dlq" {
  name          = "${var.name}-dlq"
  delay_seconds = var.delay_seconds
}

resource "aws_sqs_queue_policy" "dlq" {
  queue_url = aws_sqs_queue.dlq.id
  policy    = data.aws_iam_policy_document.dlq.json
}