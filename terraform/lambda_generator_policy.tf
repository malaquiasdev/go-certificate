resource "aws_iam_role" "lambda_generator_role" {
  name = "${var.project_name}-lambda-generator-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "lambda_generator_policy" {
  name        = format("%s-trigger-transcoder", "${var.project_name}-lambda-generator-policy")
  description = "Allow to access base resources and trigger transcoder"
  policy      = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "SomeVeryDefaultAndOpenActions",
            "Effect": "Allow",
            "Action": [
                "logs:*",
                "dynamodb:*"
            ],
            "Resource": [
                "arn:aws:logs:*:*:*",
                "arn:aws:dynamodb:${var.aws_region}:${var.aws_account_number}:table/${var.ddb_name}"
            ]
        }
    ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "lambda_generator_dynamodb" {
  policy_arn = aws_iam_policy.lambda_generator_policy.arn
  role       = aws_iam_role.lambda_generator_role.name
}