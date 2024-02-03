**Self-Signed Certificates and Secrets**:

- Use self-signed certificates for securing communication between its components/nodes.
- Managed through secrets?

**Node Authentication and TLS Bootstrapping**:

- Each node agent is required to authenticate itself to the API server using TLS certificates.
- Agent generates a certificate signing request (CSR) and submits it to the API server. The API server approves the request, providing the required certificate to the kubelet for secure communication.
