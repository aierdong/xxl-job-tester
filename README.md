# XXL-Job Tester (xxl-job-tester)

XXL-Job Tester (xxl-job-tester) is a Golang command-line tool that allows you to simulate XXL-Job Admin calls to your executor program. This can be helpful for testing your executor logic without deploying it to the XXL-Job server.

## Installation
The `xxl-job-tester` requires Golang to be installed on your system. You can download and install Golang from the official website: https://go.dev/

Once Golang is installed, you can build the xxl-job-tester executable by running the following command in the directory where the source code is located:

```Bash
go build -o xxl-job-tester
```

This will create an executable file named `xxl-job-tester` in the current directory.

## Usage

The `xxl-job-tester` accepts several command-line flags that allow you to configure the request sent to your executor:

* `-h, --host`: The hostname or IP address of your executor program (default: `localhost`)
* `-p, --port`: The port number on which your executor program is listening (default: `8080`)
* `-j, --job_handler`: The name of the job handler method in your executor program (required)
* `-P, --params`: Additional parameters to be passed to the job handler (optional)
* `-b, --block_strategy`: The blocking strategy for the job (default: `SERIAL_EXECUTION`)
* `--timeout`: The timeout for the job execution (in seconds, optional)
* `--log_id`: The log ID for the job (optional, default: `0`)
* `--log_date_time`: The log timestamp for the job (in milliseconds since epoch, optional)
* `--glue_type`: The glue type for the job (default: `BEAN`)
* `--broadcast_index`: The broadcast index for the job (optional, default: `0`)
* `--broadcast_total`: The total number of broadcasts for the job (optional, default: `1`)
* `--protocol`: The protocol used to connect to your executor (default: `http`, also supports `https`)

## Example usage:

Simulate a call to the job handler named "MyJobHandler" on your executor running on localhost:8080, passing the parameter "key1=value1":

```Bash
xxl-job-tester -j MyJobHandler -P "key1=value1"
```

This will print the following information to the console:

* The constructed URL for the request
* The formatted JSON request body
* The response status code
* The response body

> Note: The `--job_handler` flag is required and must be specified.

## Contributing

We welcome contributions to this project! If you have any suggestions or bug fixes, please feel free to create a pull request on the GitHub repository.
