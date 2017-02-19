# CTF File Server
Our simple web server that give us download links to our files on S3

## URLs

**Normal URL**: /{name}

{name} = name of the file in S3

Ex. files.cppfast.org/data.zip

**URL with folder**: /{name}?folder={foldername}

{folder} = name of the folder the file is in

Ex. files.cppfast.org/data.zip?folder=forensics

files.cppfast.org/data.zip?folder=forensics/q1/

**URL with bucket**: /{name}?bucket={bucketname}

{bucketname} = name of the bucket

## Configuring Credentials

**Before modifying using this for development, configure credentials.**

The best way to configure credentials on a development machine is to use the `~/.aws/credentials` file, which might look like:

```
[default]
aws_access_key_id = AKID1234567890
aws_secret_access_key = MY-SECRET-KEY
```

You can learn more about the credentials file from this
[blog post](http://blogs.aws.amazon.com/security/post/Tx3D6U6WSFGOK2H/A-New-and-Standardized-Way-to-Manage-Credentials-in-the-AWS-SDKs).

Alternatively, you can set the following environment variables:

```
AWS_ACCESS_KEY_ID=AKID1234567890
AWS_SECRET_ACCESS_KEY=MY-SECRET-KEY
```
