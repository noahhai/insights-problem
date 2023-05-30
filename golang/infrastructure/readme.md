# Terraform

You can use the following policy when creating the SQS queue:

```
  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "sqspolicy",
  "Statement": [
    {
      "Sid": "First",
      "Effect": "Allow",
      "Principal": "*",
      "Action": "sqs:SendMessage",
      "Resource": "${aws_sqs_queue.delinea_interview_queue.arn}"
    }
  ]
}
POLICY
```


You can read messages from your queue using the command

```
aws sqs receive-messages --queue-url <> --wait-time-seconds <>
```

Some examples of publishing to SQS from golang

https://ayada.dev/snippets/send-messages-to-aws-sqs-using-go/



You can use the following ecr repository

```
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 267294532431.dkr.ecr.us-east-1.amazonaws.com

docker build -t interview-insights .

docker tag interview-insights:latest 267294532431.dkr.ecr.us-east-1.amazonaws.com/interview-insights:latest

docker push 267294532431.dkr.ecr.us-east-1.amazonaws.com/interview-insights:latest
```
