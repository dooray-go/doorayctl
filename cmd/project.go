package cmd

import (
	"doorayctl/result"
	"log/slog"
	"os"
	"strings"

	"github.com/dooray-go/dooray/openapi/project"
	model "github.com/dooray-go/dooray/openapi/model/project"
	"github.com/spf13/cobra"
)

var (
	projectType         string
	scope               string
	state               string
	toMemberIds         string
	postWorkflowClasses string

	// post create flags
	subject     string
	content     string
	mimeType    string
	toMembers   string
	ccMembers   string
	priority    string
	workflowID  string
	milestoneID string
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Manage projects",
	Long:  `Manage Dooray projects.`,
}

var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Long:  `List all projects in Dooray.`,
	Run: func(cmd *cobra.Command, args []string) {
		log := slog.New(slog.NewTextHandler(os.Stdout, nil))

		env, err := GetEnv()
		if err != nil {
			log.Warn("Failed to get environment", "error", err)
			return
		}

		projectClient := project.NewDefaultProject()
		projectsResponse, err := projectClient.GetProjects(env.Token, projectType, scope, state)
		if err != nil {
			log.Warn("Get Projects Failed.", "error", err)
			return
		}

		err = result.PrintProjectsResult(projectsResponse)
		if err != nil {
			log.Warn("Print Projects Failed.", "error", err)
			return
		}
	},
}

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Manage project posts",
	Long:  `Manage posts in Dooray projects.`,
}

var postListCmd = &cobra.Command{
	Use:   "list [projectId]",
	Short: "List posts in a project",
	Long:  `List all posts in a specific Dooray project.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log := slog.New(slog.NewTextHandler(os.Stdout, nil))

		env, err := GetEnv()
		if err != nil {
			log.Warn("Failed to get environment", "error", err)
			return
		}

		projectId := args[0]

		projectClient := project.NewDefaultProject()
		postsResponse, err := projectClient.GetPosts(env.Token, projectId, toMemberIds, postWorkflowClasses)
		if err != nil {
			log.Warn("Get Posts Failed.", "error", err)
			return
		}

		err = result.PrintPostsResult(postsResponse)
		if err != nil {
			log.Warn("Print Posts Failed.", "error", err)
			return
		}
	},
}

var postCreateCmd = &cobra.Command{
	Use:   "create [projectId]",
	Short: "Create a new post in a project",
	Long:  `Create a new post in a specific Dooray project.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log := slog.New(slog.NewTextHandler(os.Stdout, nil))

		env, err := GetEnv()
		if err != nil {
			log.Warn("Failed to get environment", "error", err)
			return
		}

		projectId := args[0]

		// Validate required fields
		if subject == "" {
			log.Warn("Subject is required")
			return
		}
		if content == "" {
			log.Warn("Content is required")
			return
		}

		// Set default mime type if not specified
		if mimeType == "" {
			mimeType = "text/plain"
		}

		// Build post request
		postRequest := model.PostRequest{
			Subject: subject,
			Body: model.PostBody{
				MimeType: mimeType,
				Content:  content,
			},
		}

		// Add users if specified
		if toMembers != "" || ccMembers != "" {
			users := &model.PostUsers{}

			if toMembers != "" {
				memberIds := strings.Split(toMembers, ",")
				for _, memberId := range memberIds {
					memberId = strings.TrimSpace(memberId)
					if memberId != "" {
						users.To = append(users.To, model.PostRecipient{
							Type: "member",
							Member: &model.PostMember{
								OrganizationMemberID: memberId,
							},
						})
					}
				}
			}

			if ccMembers != "" {
				memberIds := strings.Split(ccMembers, ",")
				for _, memberId := range memberIds {
					memberId = strings.TrimSpace(memberId)
					if memberId != "" {
						users.Cc = append(users.Cc, model.PostRecipient{
							Type: "member",
							Member: &model.PostMember{
								OrganizationMemberID: memberId,
							},
						})
					}
				}
			}

			postRequest.Users = users
		}

		// Add optional fields
		if priority != "" {
			postRequest.Priority = priority
		}
		if workflowID != "" {
			postRequest.WorkflowID = workflowID
		}
		if milestoneID != "" {
			postRequest.MilestoneID = milestoneID
		}

		projectClient := project.NewDefaultProject()
		postResponse, err := projectClient.CreatePost(env.Token, projectId, postRequest)
		if err != nil {
			log.Warn("Create Post Failed.", "error", err)
			return
		}

		log.Info("Post created successfully", "postId", postResponse.Result.ID)
	},
}

func init() {
	projectListCmd.Flags().StringVarP(&projectType, "type", "t", "", "Project type (e.g., project, blog)")
	projectListCmd.Flags().StringVarP(&scope, "scope", "s", "", "Project scope (e.g., public, private)")
	projectListCmd.Flags().StringVar(&state, "state", "", "Project state (e.g., active, archived)")

	postListCmd.Flags().StringVar(&toMemberIds, "to-members", "", "Filter by member IDs (comma-separated)")
	postListCmd.Flags().StringVar(&postWorkflowClasses, "workflow-classes", "", "Filter by workflow classes (e.g., registered,working,closed)")

	postCreateCmd.Flags().StringVarP(&subject, "subject", "s", "", "Post subject (required)")
	postCreateCmd.Flags().StringVarP(&content, "content", "c", "", "Post content (required)")
	postCreateCmd.Flags().StringVarP(&mimeType, "mime-type", "m", "text/plain", "Content MIME type (default: text/plain)")
	postCreateCmd.Flags().StringVar(&toMembers, "to", "", "Assignee member IDs (comma-separated)")
	postCreateCmd.Flags().StringVar(&ccMembers, "cc", "", "CC member IDs (comma-separated)")
	postCreateCmd.Flags().StringVarP(&priority, "priority", "p", "", "Priority (urgent | high | normal | low)")
	postCreateCmd.Flags().StringVar(&workflowID, "workflow-id", "", "Workflow ID")
	postCreateCmd.Flags().StringVar(&milestoneID, "milestone-id", "", "Milestone ID")
	postCreateCmd.MarkFlagRequired("subject")
	postCreateCmd.MarkFlagRequired("content")

	postCmd.AddCommand(postListCmd, postCreateCmd)
	projectCmd.AddCommand(projectListCmd, postCmd)
	rootCmd.AddCommand(projectCmd)
}
