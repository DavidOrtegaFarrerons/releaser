# Release Handler

A CLI tool that integrates Jira issues and Azure DevOps pull requests. It fetches relevant data, merges it into a table, and opens related links in a browser for easy navigation.

---

## 🚀 Features

- Fetches Jira issues and Azure DevOps pull requests.
- Displays a structured table of tickets and PRs.
- Supports opening Jira issues and Azure PRs in the browser.
- Uses **Viper** for configuration management.
- Supports multiple platforms (Windows, macOS, Linux).

---

## 🛠 Installation

### Prerequisites
- Go (>= 1.18)
- Jira API access
- Azure DevOps API access

### Build from Source
```sh
git clone https://github.com/DavidOrtegaFarrerons/releaser.git
cd releaser
go mod tidy
go build -o releaser main.go
```

## ⚙️ Configuration

The project uses **Viper** for configuration management. You can configure your Jira and Azure DevOps access in the `.release-handler.yaml` file that is generated automatically in your $HOME once you run the program.

### Sample `config.yaml`:

```yaml
"azure_key": ""
"azure_organization": ""
"azure_project": ""
"azure_repository_id": ""
"jira_domain": ""
"jira_email": ""
"jira_jql": ""
"jira_key": ""
"last_opened_url": ""
"ticket_prefix": ""
```

### Generate and display the tickets/PRs table
```
./releaser table
```

### Generate a tag for release (format: release-YYYY-MM-DD)
```
./releaser tag [rls|release]
```

### Generate a tag for production ((format: prod-YYYY-MM-DD)
```
./releaser tag [prd|prod|production]


🔮 Future Improvements
----------------------

Planned enhancements for upcoming versions:

-   **Multi-repo Support**: Handle pull requests from multiple repositories like bitbucket, github and gitlab

-   **Tag creation**: Support for creating custom tags for the releases --> This is already implemented 

-   **Enhanced Configuration**: GUI configurator and validation to make setup of the tool easier

-   **Ticket Status Update**: Once a PR is detected as merged, you will have the option to update the ticket status

