# Baremetrics SMS

A simple Heroku deployable app which will SMS you and your team with updates on your MRR and growth.

Requirements:
* a heroku account
* a twilio account
* a baremetrics account

## Installation

### Keys

You'll need to collect the following keys:

1. ``TWILIO_SID`` & ``TWILIO_AUTH`` - Get your Twilio SID and Auth token from here: https://www.twilio.com/user/account/settings
2. ``TWILIO_FROM`` - Get an incoming number from here: https://www.twilio.com/user/account/phone-numbers/incoming
3. ``BAREMETRICS_COOKIE`` - to get this, login to Baremetrics, then open the Chrome Developer Console. Use the network tab, select a request to baremetrics.com and inspect the request headers. There will be a field called cookie; copy the value.

### Deploy

Click this button, then fill in the details from above:

[![Deploy to Heroku](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

### Configure

1. Manage the app (aka open the Heroku Dashboard for the app you just created)
2. Click on the ``Heroku Scheduler`` addon
3. Add a schedule
4. Select the time you want, enter ``/app/bin/baremetrics-sms`` as the command

### Test

Run:
```
heroku run /app/bin/baremetrics-sms -a <your-heroku-app-name>
```

You should recieve a text :)

## License

See [here](LICENSE.txt)