# mikrotik_dns_record

Creates a DNS record on the mikrotik device

## Example Usage

```hcl
resource "mikrotik_dns_record" "record" {
  name = "example.domain.com"
  address = "192.168.88.1"
  ttl = 300
}
```

## Argument Reference
* name - (Required) The name of the DNS hostname to be created
* address - (Required) The A record to be returend from the DNS hostname
* ttl - (Optional) The ttl of the DNS record.
* comment - (Optional) The comment text associated with the DNS record.

## Attributes Reference

## Import Reference

```bash
terraform import mikrotik_dns_record.record example.domain.com
```
