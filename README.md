# goremindme

Tool, that sends reminder through OS notifications for tasks.

```sh
$ ./bin/app
  -e string
    	Time between consecutive reminders, e.g. 2m, 5h, 1d
  -m string
    	Message to be displayed as reminder, e.g. 'Stay focused'
  -p string
    	Period for the reminder to run, e.g. 10m, 5m, 1h, 1d (Required)
  -t string
    	Task name, e.g. 'Task notifier' (Required)
```