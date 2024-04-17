data "aws_iam_policy_document" "main" {
  statement {
    effect    = "Allow"
    resources = [aws_sqs_queue.main.arn]
    actions = [
      "sqs:ChangeMessageVisibility",
      "sqs:DeleteMessage",
      "sqs:GetQueueAttributes",
      "sqs:GetQueueUrl",
      "sqs:ListQueueTags",
      "sqs:ReceiveMessage",
      "sqs:SendMessage",
    ]
    principals {
      type        = "AWS"
      identifiers = ["*"]
    }
  }
}

data "aws_iam_policy_document" "dlq" {
  statement {
    effect    = "Allow"
    resources = [aws_sqs_queue.dlq.arn]
    actions = [
      "sqs:ChangeMessageVisibility",
      "sqs:DeleteMessage",
      "sqs:GetQueueAttributes",
      "sqs:GetQueueUrl",
      "sqs:ListQueueTags",
      "sqs:ReceiveMessage",
      "sqs:SendMessage",
    ]
    principals {
      type        = "AWS"
      identifiers = ["*"]
    }
  }
}