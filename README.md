<p align="center">
  <h1 align="center">clighl</h1>
  <p align="center">A cross-platform CLI for <a href="https://www.gohighlevel.com/">Go HighLevel</a> — control your entire CRM from the terminal.</p>
</p>

<p align="center">
  <a href="#installation">Installation</a> &nbsp;&bull;&nbsp;
  <a href="#quick-start">Quick Start</a> &nbsp;&bull;&nbsp;
  <a href="#commands">Commands</a> &nbsp;&bull;&nbsp;
  <a href="#api-coverage">API Coverage</a> &nbsp;&bull;&nbsp;
  <a href="https://marketplace.gohighlevel.com/docs/">GHL API Docs</a>
</p>

---

Built on the [Go HighLevel API v2](https://marketplace.gohighlevel.com/docs/). Covers contacts, pipelines, opportunities, calendars, conversations, notes, tags, blogs, social media, email templates, and payments — **52 commands** across **13 resource groups**.

Designed for developers, shell scripts, CI/CD pipelines, and AI agents. Human-readable tables by default, `--json` for machine consumption.

## Prerequisites

- A [Go HighLevel](https://www.gohighlevel.com/) account
- A **Location ID** (sub-account ID) and **Private Integration Token (PIT)**
  - Create a PIT under **Settings > Integrations > Private Integrations** in your GHL sub-account
  - See the [GHL Private Integrations guide](https://marketplace.gohighlevel.com/docs/Authorization/privateIntegrations/) for details

## Installation

### Pre-built Binaries

Download the latest binary for your platform from [Releases](https://github.com/altusmusic/clighl/releases):

<details>
<summary><strong>macOS</strong></summary>

```bash
# Apple Silicon (M1/M2/M3/M4)
curl -Lo clighl.tar.gz https://github.com/altusmusic/clighl/releases/latest/download/clighl_Darwin_arm64.tar.gz
tar xzf clighl.tar.gz && sudo mv clighl /usr/local/bin/

# Intel
curl -Lo clighl.tar.gz https://github.com/altusmusic/clighl/releases/latest/download/clighl_Darwin_amd64.tar.gz
tar xzf clighl.tar.gz && sudo mv clighl /usr/local/bin/
```
</details>

<details>
<summary><strong>Linux</strong></summary>

```bash
# amd64
curl -Lo clighl.tar.gz https://github.com/altusmusic/clighl/releases/latest/download/clighl_Linux_amd64.tar.gz
tar xzf clighl.tar.gz && sudo mv clighl /usr/local/bin/

# arm64
curl -Lo clighl.tar.gz https://github.com/altusmusic/clighl/releases/latest/download/clighl_Linux_arm64.tar.gz
tar xzf clighl.tar.gz && sudo mv clighl /usr/local/bin/
```
</details>

<details>
<summary><strong>Windows</strong></summary>

```powershell
Invoke-WebRequest -Uri https://github.com/altusmusic/clighl/releases/latest/download/clighl_Windows_amd64.zip -OutFile clighl.zip
Expand-Archive clighl.zip -DestinationPath .
Move-Item clighl.exe C:\Windows\System32\
```
</details>

### From Source

Requires [Go 1.21+](https://go.dev/dl/):

```bash
go install github.com/altusmusic/clighl@latest
```

Or clone and build:

```bash
git clone https://github.com/altusmusic/clighl.git
cd clighl
make build     # binary at ./clighl
make install   # moves to $GOPATH/bin
```

## Quick Start

```bash
# 1. Authenticate (interactive — prompts for Location ID and PIT token)
clighl auth

# 2. Check your location
clighl location info

# 3. Search contacts
clighl contacts search "John"

# 4. Move a contact through your pipeline
clighl opportunities move "John Smith" --pipeline "Sales" --stage "Qualified"

# 5. Book a calendar appointment
clighl cal slots "Discovery Call" --days 7
clighl cal book --calendar "Discovery Call" --contact "John" \
  --slot "2026-04-01T10:00:00-04:00" --timezone "America/New_York"

# 6. Send a message
clighl conv send --contact "John" --type SMS --message "Hey John, confirming our call tomorrow"

# 7. Tag the contact
clighl tags add --contact "John" --tags "qualified,booked"
```

> **Tip:** You don't need IDs. `clighl` resolves contacts, pipelines, stages, and calendars by name. If multiple contacts match, you get an interactive picker.

---

## Commands

### Authentication

```
clighl auth                    Interactive credential setup
clighl auth status             Show current config (token masked)
clighl auth logout             Remove stored credentials
```

Credentials are validated against the API before saving. Stored in `~/.clighl/config.yaml` with `0600` permissions.

### Location

```
clighl location info           Location name, email, phone, timezone, address
clighl location fields         List custom field definitions (id, name, key, type)
```

> GHL API: [Get Location](https://marketplace.gohighlevel.com/docs/ghl/locations/get-location/), [Get Custom Fields](https://marketplace.gohighlevel.com/docs/ghl/locations/get-custom-fields/)

### Contacts

```
clighl contacts search <query>           Search by name, email, or phone
clighl contacts get <id>                 Full contact details
clighl contacts list                     Paginated list (--limit, --page)
clighl contacts create                   Interactive or via flags
clighl contacts update <name>            Update fields (--email, --phone, etc.)
clighl contacts upsert                   Create or update by email/phone match
clighl contacts tasks <name>             List tasks assigned to a contact
```

**Create/update flags:** `--first-name`, `--last-name`, `--email`, `--phone`, `--company`

> GHL API: [Search Contacts](https://marketplace.gohighlevel.com/docs/ghl/contacts/search-contacts-advanced/), [Get Contact](https://marketplace.gohighlevel.com/docs/ghl/contacts/get-contact/), [Create Contact](https://marketplace.gohighlevel.com/docs/ghl/contacts/create-contact/), [Update Contact](https://marketplace.gohighlevel.com/docs/ghl/contacts/update-contact/), [Upsert Contact](https://marketplace.gohighlevel.com/docs/ghl/contacts/upsert-contact/), [Get Tasks](https://marketplace.gohighlevel.com/docs/ghl/contacts/get-all-tasks/)

### Pipelines & Opportunities

```
clighl pipelines list                    All pipelines with stages (shown as flow)
clighl pipelines get <id-or-name>        Pipeline detail with stage IDs

clighl opportunities list                List opportunities (--pipeline filter)
clighl opportunities get <id>            Opportunity detail
clighl opportunities create              Create (--contact, --pipeline, --stage)
clighl opportunities move <contact>      Move contact to pipeline stage
```

The `move` command is the flagship — it resolves contact, pipeline, and stage by name, then creates a new opportunity or updates the existing one:

```bash
clighl opp move "Dan" --pipeline "Leads Ads" --stage "Contacted" --value 5000
```

> GHL API: [Get Pipelines](https://marketplace.gohighlevel.com/docs/ghl/opportunities/get-pipelines/), [Search Opportunities](https://marketplace.gohighlevel.com/docs/ghl/opportunities/search-opportunity/), [Create Opportunity](https://marketplace.gohighlevel.com/docs/ghl/opportunities/create-opportunity/), [Update Opportunity](https://marketplace.gohighlevel.com/docs/ghl/opportunities/update-opportunity/)

**Alias:** `clighl opp` = `clighl opportunities`

### Calendars

```
clighl cal list                          All calendars (name, duration, active)
clighl cal get <id-or-name>              Calendar detail with slot config
clighl cal slots <calendar>              Available time slots (--days, --timezone)
clighl cal events                        List events (--calendar, --start, --end)
clighl cal book                          Book an appointment
clighl cal appointment <id>              Appointment detail
clighl cal notes <appointment-id>        Appointment notes
clighl cal cancel <id>                   Cancel an appointment
```

**Booking workflow:**
```bash
clighl cal list                                          # find calendar name
clighl cal slots "Sales Call" --days 7 --timezone "US/Eastern"  # find a slot
clighl cal book --calendar "Sales Call" --contact "Dan" \
  --slot "2026-04-01T10:00:00-04:00" \
  --timezone "America/New_York" \
  --notes "Initial consultation"
```

> GHL API: [Calendars](https://marketplace.gohighlevel.com/docs/ghl/calendars/get-calendars/), [Free Slots](https://marketplace.gohighlevel.com/docs/ghl/calendars/get-free-slots/), [Create Appointment](https://marketplace.gohighlevel.com/docs/ghl/calendars/create-appointment/), [Calendar Events](https://marketplace.gohighlevel.com/docs/ghl/calendars/get-calendar-events/)

**Alias:** `clighl cal` = `clighl calendars`

### Conversations & Messaging

```
clighl conv search                       Search conversations (--contact, --query)
clighl conv messages <conversation-id>   Message history (--limit, --after)
clighl conv send                         Send SMS, Email, or WhatsApp
```

```bash
# Find conversations for a contact
clighl conv search --contact "Dan"

# Read messages
clighl conv messages k875FrRiRO5ylPsgPRJp --limit 10

# Send an SMS
clighl conv send --contact "Dan" --type SMS --message "Following up on our call"

# Send an email by contact name or email
clighl conv send --contact "dan@test.com" --type Email \
  --subject "Proposal attached" \
  --message "<h1>Hi Dan</h1><p>Here's the proposal.</p>"

# Send an email by exact contact ID
clighl conv send --contact "JY81tMqvDZJG4EiGDujA" --type Email \
  --subject "Quick test" \
  --message "<p>This is a live test email from clighl.</p>"
```

**Email note:** for `--type Email`, pass the email body with `--message`. HTML content is supported, and exact contact IDs now resolve correctly for non-interactive use.

> GHL API: [Search Conversations](https://marketplace.gohighlevel.com/docs/ghl/conversations/search-conversation/), [Get Messages](https://marketplace.gohighlevel.com/docs/ghl/conversations/get-messages/), [Send Message](https://marketplace.gohighlevel.com/docs/ghl/conversations/send-a-new-message/)

**Alias:** `clighl conv` = `clighl conversations`

### Notes

```
clighl notes list --contact <name>               List all notes for a contact
clighl notes add --contact <name> --body "..."   Add a note
clighl notes delete <note-id> --contact <name>   Delete a note
```

> GHL API: [Contact Notes](https://marketplace.gohighlevel.com/docs/ghl/contacts/)

### Tags

```
clighl tags list                                          All location tags
clighl tags add --contact <name> --tags "vip,hot-lead"    Add tags to contact
clighl tags remove --contact <name> --tags "old-lead"     Remove tags
```

> GHL API: [Add Tags](https://marketplace.gohighlevel.com/docs/ghl/contacts/add-tags/), [Remove Tags](https://marketplace.gohighlevel.com/docs/ghl/contacts/remove-tags/), [Get Tags](https://marketplace.gohighlevel.com/docs/ghl/locations/get-location-tags/)

### Blogs

```
clighl blogs list                        All blog sites
clighl blogs posts <blog-id>             Posts for a blog (--limit, --offset)
clighl blogs create                      Create a post (--blog, --title, --html, --status)
clighl blogs update <post-id>            Update a post (--blog, --title, --status)
clighl blogs slug-check                  Check URL slug availability (--blog, --slug)
clighl blogs authors                     List blog authors
clighl blogs categories                  List blog categories
```

> GHL API: [Blogs](https://marketplace.gohighlevel.com/docs/ghl/blogs/)

### Social Media

```
clighl social accounts                   Connected social accounts
clighl social stats                      Account statistics (--account)
clighl social posts                      List posts (--account, --limit, --page)
clighl social get <post-id>              Single post detail
clighl social create                     Create post (--accounts, --content, --schedule)
clighl social update <post-id>           Update post (--content, --schedule)
```

> GHL API: [Social Media Posting](https://marketplace.gohighlevel.com/docs/ghl/socialmediaposting/)

### Email Templates

```
clighl emails list                       All email templates
clighl emails create                     Create template (--name, --subject, --html)
```

> GHL API: [Email Templates](https://marketplace.gohighlevel.com/docs/ghl/emails/)

### Payments

```
clighl pay order <id>                    Order details
clighl pay transactions                  List transactions (--contact, --limit, --page)
```

> GHL API: [Payments](https://marketplace.gohighlevel.com/docs/ghl/payments/)

**Alias:** `clighl pay` = `clighl payments`

---

## Output Formats

Every command supports two output modes:

```bash
# Human-readable tables (default)
clighl contacts list

# JSON for scripting, piping, and automation
clighl contacts list --json
clighl contacts list --json | jq '.[].email'
clighl pipelines list --json | jq '.[] | {name, id}'
```

Errors always go to stderr, so JSON piping works cleanly.

## How It Works

### Architecture

```
~/.clighl/config.yaml          Credentials (location_id, access_token)
        |
    clighl CLI                  Cobra commands → internal/api → GHL API v2
        |
https://services.leadconnectorhq.com
```

All requests include:
- `Authorization: Bearer {token}` — your PIT or OAuth token
- `Version: 2021-07-28` — required by GHL API v2
- Built-in rate limiting (10 req/s, 100 burst) with automatic retry on 429

### Contact-Centric Design

Most CRM operations need a contact ID. `clighl` resolves names automatically:

1. Searches by name/email/phone
2. **1 match** — uses it directly
3. **Multiple matches** — interactive picker (TTY) or error with guidance (non-TTY/scripts)
4. **No matches** — error with search suggestions

This works everywhere: `--contact "Dan"`, `contacts update "Dan"`, `contacts tasks "Dan"`, `opp move "Dan"`, etc.

### Name Resolution

Pipelines, stages, and calendars are also resolved by name (case-insensitive):

```bash
# No IDs needed anywhere
clighl opp move "Dan" --pipeline "Sales" --stage "Lead"
clighl cal slots "Discovery Call"
clighl cal book --calendar "Onboarding" --contact "Dan" --slot "..."
```

## Configuration

**Config file:** `~/.clighl/config.yaml`

```yaml
location_id: your-location-id
access_token: pit-your-token
api_version: "2021-07-28"
```

**Environment variables** override the config file (useful for CI/CD and scripting):

```bash
export CLIGHL_LOCATION_ID="your-location-id"
export CLIGHL_ACCESS_TOKEN="pit-your-token"

# Now any command works without a config file
clighl contacts list --json
```

## API Coverage

All [popular GHL API v2 endpoints](https://marketplace.gohighlevel.com/docs/) are covered:

| Group | Commands | GHL API Reference |
|-------|----------|-------------------|
| Auth | auth, status, logout | [Authorization](https://marketplace.gohighlevel.com/docs/Authorization/) |
| Location | info, fields | [Locations](https://marketplace.gohighlevel.com/docs/ghl/locations/) |
| Contacts | search, get, list, create, update, upsert, tasks | [Contacts](https://marketplace.gohighlevel.com/docs/ghl/contacts/) |
| Pipelines | list, get | [Opportunities](https://marketplace.gohighlevel.com/docs/ghl/opportunities/) |
| Opportunities | list, get, create, move | [Opportunities](https://marketplace.gohighlevel.com/docs/ghl/opportunities/) |
| Calendars | list, get, slots, events, book, appointment, notes, cancel | [Calendars](https://marketplace.gohighlevel.com/docs/ghl/calendars/) |
| Conversations | search, messages, send | [Conversations](https://marketplace.gohighlevel.com/docs/ghl/conversations/) |
| Notes | list, add, delete | [Contacts](https://marketplace.gohighlevel.com/docs/ghl/contacts/) |
| Tags | list, add, remove | [Contacts](https://marketplace.gohighlevel.com/docs/ghl/contacts/) / [Locations](https://marketplace.gohighlevel.com/docs/ghl/locations/) |
| Blogs | list, posts, create, update, slug-check, authors, categories | [Blogs](https://marketplace.gohighlevel.com/docs/ghl/blogs/) |
| Social Media | accounts, stats, posts, get, create, update | [Social Media](https://marketplace.gohighlevel.com/docs/ghl/socialmediaposting/) |
| Email Templates | list, create | [Emails](https://marketplace.gohighlevel.com/docs/ghl/emails/) |
| Payments | order, transactions | [Payments](https://marketplace.gohighlevel.com/docs/ghl/payments/) |

> **Note:** Some endpoints (Blogs, Social Media, Emails, Payments) require specific [OAuth scopes](https://marketplace.gohighlevel.com/docs/Authorization/OAuth2.0/) or plan-level features. If you get 403/404 errors, check your GHL app permissions under **Settings > Integrations**.

## Development

```bash
make build        # Build binary
make test         # Run tests
make install      # Install to PATH
make lint         # Run golangci-lint
make clean        # Remove build artifacts

# Cross-compile verification
GOOS=windows go build .
GOOS=linux go build .
GOOS=darwin go build .
```

### Project Structure

```
clighl/
├── main.go                      # Entry point
├── cmd/                         # Cobra commands (one file per command)
│   ├── root.go                  # Global flags (--json, --verbose)
│   ├── auth.go                  # Auth group
│   ├── contacts.go              # Contacts group + shared helpers
│   ├── calendars.go             # Calendars group
│   └── ...                      # 60+ command files
├── internal/
│   ├── api/                     # GHL API client (HTTP, rate limiting, auth)
│   ├── config/                  # Config load/save (~/.clighl/config.yaml)
│   ├── models/                  # Request/response structs with JSON tags
│   ├── resolver/                # Name → ID resolution with disambiguation
│   └── output/                  # Table + JSON formatters
├── .goreleaser.yaml             # Cross-platform release builds
├── Makefile                     # Build shortcuts
└── README.md
```

## Contributing

1. Fork the repo
2. Create a feature branch (`git checkout -b feature/workflows`)
3. Add your commands following the existing pattern in `cmd/`
4. Ensure `go vet ./...` passes and cross-platform builds work
5. Open a PR

## License

[MIT](LICENSE)
