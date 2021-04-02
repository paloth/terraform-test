# Introduction to Terraform test command

## How to test in terraform

Test is a new feature in 0.15 version (Still experimental)

It allows you to perform integration test without a third-party like Terratest.  
The Test will deploy the infrastructure on an AWS account (You need to setup credentials). 
It will perform `plan`, `apply`, perform tests with the collected data, `destroy`

## File structure

In your repository create a `test` or `tests` folder.  
Inside this folder you can create one or more folders that will contains your tests cases.

```
├──tests
│  ├──test_case_1
│  │  └──test_case_1.tf
│  └──test_case_2
│     └──test_case_2.tf
├──.gitignore
├──.terraform.lock.hcl
├──LICENSE
├──README.md
├──backend.tf
├──provider.tf
└──version.tf
```

Only the file `.tf` must start with `test_`.  
It does not matter if the folder's name contains `test`

## Test_file.tf

The test files must use a specific provider plus the provider you want to use in your terraform stack. 

```hcl
terraform {
  required_providers {
    test = {
      source = "terraform.io/builtin/test"
    }
  }
}
```

Import your module inside your test case with the standard hcl block

```hcl
module "main" {
  source = "PATH_TO_YOUR_MODULE"
}
```

In the test file, the principal resource to use is test_assertion:

```hcl
resource "test_assertions" "my_asssertion" {}
```

This resource must contain the `component` parameter.  
`component` is an unique identifier for the assertion resource in the test results.

Output value to point to the failed test will be displayed with the format `folder_name.component.[equal or check block]_name`

Example:

```
bucket.bucket_region_from_aws.region
```

Block `equal` or `check` are used inside to define the tests

```hcl
equal "my_equal" {
    description = ""
    got         = The value deployed
    want        = The correct value i need
  }

  check "my_check" {
    description = ""
    condition   = Use a fonction that return true or false like can() 
  }

```

With these block you can check the module's `outputs`.  
You can also check the value of a deployed resource by fetching the information with a `data_source` block.

### Example of test outputs:

Failed:

Check Test:

```
─── Failed: test_case_folder.component.test_name (Test description) ───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
condition failed
```

Equal Test:

```
─── Failed: test_case_folder.component.test_name (Test description) ───────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
wrong value
    got:  "value1"
    want: "value2"
```

Success:

```
Success! All of the test assertions passed.
```
