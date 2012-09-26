#ringgo
ringgo is a tool, written in Go, to act as a bridge between Gomez's URL "Alert Destination" and [RingRing](https://github.com/darrendao/ringring), an on call/calendaring tool. Each on call list can be queried for the email/SMS number(s) using the "Smart Contacts" feature. When ringgo gets an HTTP POST from Gomez to http://ringgo.blah.com/gomezalerts/<list_name>, ringgo will take the value for list_name, query RingRing, and pass along the alert to the appropriate folks.

##Requirements
* the Go programming language 1.0.x

##Getting Started
ringgo should sit on a server/vm outside of your data center that is both available to Gomez and has access to RingRing. The idea is that even if your datacenter/email service/RingRing itself is down or otherwise unavailable, Gomez still has a place to send the alert. If the REST call to RingRing fails from ringgo, you can set a list of folks to get notified that something is wrong, so you aren't left in the dark if something falls offline.

* enter your RingRing URL for the ringring variable
* enter your FROM email address for the email_from variable
* enter your smtp server (the default is smtp.google.com:) for the smtp_server variable
* enter your smtp login for the smtp_login variable
* enter your smtp login's password for the smtp_passwd variable
* enter your smtp host name for the smtp_host variable

## TODO
* all of the FIXMEs

##Getting Started
TODO: make this section less worse.

