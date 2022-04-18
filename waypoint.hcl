project = "tfc-auto-apply"

app "auto-apply-lambda" {

  config {
    env = {
      TFE_TOKEN = dynamic("aws_ssm", {
        path = "tfe_token"
      })
    }
  }

  build {
    use "pack" {}
    registry {
      use "aws-ecr" {
        region     = "us-east-1"
        repository = "auto-apply-lambda"
        tag        = gitrefpretty()
      }
    }
  }

  deploy {
    use "aws-lambda" {
      region = "us-east-1"
    }
  }

}

variable "AWS_SECRET_ACCESS_KEY" {}
variable "AWS_ACCESS_KEY_ID" {}
