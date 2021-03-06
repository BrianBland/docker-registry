Docker-Registry Configuration
=============================

This document describes the registry configuration model and how to specify a custom configuration with a configuration file and/or environment variables.

Semantic-ish Versioning
-----------------------

The configuration file is designed with versioning in mind, such that most upgrades will not require a change in configuration files, and such that configuration files can be "upgraded" from one version to another.

The version is specified as a string of the form `MajorVersion.MinorVersion`, where MajorVersion and MinorVersion are both non-negative integer values. Much like [semantic versioning](http://semver.org/), minor version increases denote inherently backwards-compatible changes, such as the addition of optional fields, whereas major version increases denote a restructuring, such as renaming fields or adding required fields. Because of the explicit version definition in the configuration file, it should be possible to parse old configuration files and port them to the current configuration version, although this is not guaranteed for all future versions.

File Structure (as of Version 0.1)
------------------------------------

The configuration structure is defined by the `Configuration` struct in `configuration.go`, and is best described by the following two examples:

```yaml
version: 0.1
loglevel: info
storage:
  s3:
    region: us-east-1
    bucket: my-bucket
    rootpath: /registry
    encrypt: true
    secure: false
    accesskey: SAMPLEACCESSKEY
    secretkey: SUPERSECRET
    host: ~
    port: ~
```

```yaml
version: 0.1
loglevel: debug
storage: inmemory
```

### version
The version is expected to remain a top-level field, as to allow for a consistent version check before parsing the remainder of the configuration file.

### loglevel
This specifies the log level of the registry.

Supported values:
* `error`
* `warn`
* `info`
* `debug`

### storage
This specifies the storage driver, and may be provided either as a string (only the driver type) or as a driver name with a parameters map, as seen in the first example above.

The parameters map will be passed into the factory constructor of the given storage driver type.

### Notes

All keys in the configuration file **must** be provided as a string of lowercase letters and numbers only, and values must be string-like (booleans and numerical values are fine to parse as strings).

Environment Variables
---------------------

To support the workflow of running a docker registry from a standard container without having to modify configuration files, the registry configuration also supports environment variables for overriding fields.

Any configuration field other than version can be replaced by providing an environment variable of the following form: `REGISTRY_<uppercase key>[_<uppercase key>]...`.

For example, to change the loglevel to `error`, one can provide `REGISTRY_LOGLEVEL=error`, and to change the s3 storage driver's region parameter to `us-west-1`, one can provide `REGISTRY_STORAGE_S3_LOGLEVEL=us-west-1`.

### Notes
If an environment variable changes a map value into a string, such as replacing the storage driver type with `REGISTRY_STORAGE=filesystem`, then all sub-fields will be erased. As such, changing the storage type will remove all parameters related to the old storage type.

By restricting all keys in the configuration file to lowercase letters and numbers, we can avoid any potential environment variable mapping ambiguity.
