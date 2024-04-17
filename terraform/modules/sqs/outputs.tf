output "main_arn" {
  value = aws_sqs_queue.main.arn
}

output "main_url" {
  value = aws_sqs_queue.main.url
}

output "dlq_arn" {
  value = aws_sqs_queue.dlq.arn
}