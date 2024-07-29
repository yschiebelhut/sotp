# SOTP - The **S**imple **OTP** Client
> **Notice:**
> This project is provided *as is*.
> I do not give any promises in regard to security nor do I take any responsibility on errors that might be experienced during use.
>
> This version was written and tested on macOS.
> Probably also works well on Linux.
> Might also work on Windows.

At the moment, this is just a small hobbyist project out of annoyance that I have to use my phone every time I need to access some services while I could also use my PC to generate the required OTP codes.
Maybe it will grow in the future.
The basic set of functions provided at the moment is:
- adding/removing OTP keys
- generating OTPs
- command completion and integrated help
- automatic completion of key names
- copying the generated OTP to the clipboard for easy access

## Notes on Setup
This project uses a YAML file as its database.
The file's location is hard coded to `~/.config/sotp/db.yaml`.
The syntax of the file looks something like:

```yaml
secrets:
  <secret name>: <secret key>
```

Where *secret name* is the choosen name of the OTP key and *secret key* corresponds to the shared secret which is typically encoded into the QR code when using authenticator apps.
At the moment, there is no functionality to bootstrap this config location and initial file.
Although, it is possible to later add/remove keys via *sotp* without a need to manually edit the config file.
