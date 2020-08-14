# Drone.io CI plugin for Twilio

## Usage

```yml
- name: twilio
  image: royzwambag/twilio
  settings:
    account_sid: AC1234abcdef
    auth_token: secret
    from_number: +12512205274
    to_number: +12512205274
    body: Your drone build is ready!
```

Required settings:
* account_sid
* auth_token
* from_number
* to_number
* body

---

Check out the go twilio api at https://github.com/sfreiberg/gotwilio
