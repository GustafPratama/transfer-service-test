# Architecture Overview

This system implements a simple Transfer Service with clean separation of concerns:

- **Handler → Service → Repository** layers for maintainability.
- **Outbox Pattern** ensures reliable event publishing.
- **PostgreSQL** used as main data store.
- **Frontend (React + TypeScript)** provides UI for transfer and temperature alert monitoring.

### Workflow
1. User creates a transfer (POST /transfers)
2. Supervisor can Accept or Complete
3. When Completed, stock updates and event saved in outbox
4. `/dev/flush-outbox` exports events to JSON

