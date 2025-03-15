openssl req -new -newkey rsa:2048 -nodes -keyout server-key.pem -out server-csr.pem -config <(
cat <<-EOF
[req]
default_bits       = 2048
default_md         = sha256
prompt             = no
distinguished_name = dn
req_extensions     = req_ext

[dn]
CN = localhost

[req_ext]
subjectAltName = @alt_names

[alt_names]
DNS.1 = java-service
DNS.2 = go-service
DNS.3 = localhost
EOF
)
