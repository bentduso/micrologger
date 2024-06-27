/*
Package microlog provides a lightweight logging utility for microservices.

To create a logger, use the New function, specifying the desired logging
level and optional configuration functions. Messages of lesser criticality
won't be logged based on the configured threshold.

Sharing the logger is the caller's responsibility.

Once instantiated, you can use the logger to output messages based on
different logging levels:
  - Trace: Detailed tracing of code execution, capturing fine-grained details.
  - Debug: Diagnostic information useful for troubleshooting and administration.
  - Info: Updates and milestones about service operation.
  - Warn: Signals potential issues impacting system functionality.
  - Error: Recoverable errors requiring attention but not halting execution.
  - Fatal: Critical errors mandating immediate application shutdown.
*/
package microlog
