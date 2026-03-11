package handlers

import (
	"net/http"

	"github.com/tedewaard/folio/src/models"
	"github.com/tedewaard/folio/src/templates"
)

// HandleHome renders the home page with all projects and timeline data
func HandleHome(w http.ResponseWriter, r *http.Request) {
	// Work projects data
	workProjects := []models.Project{
		{
			Title:       "On-prem Kubernetes (RKE2)",
			Description: "Re-architected our 5 year old on-prem Kubernetes clusters, enabling simplified disaster recovery, continuous updates and automated cluster deployment reducing our RTO by over 90%.",
			Technologies: []models.Technology{
				{Name: "Kubernetes (RKE2)", Color: "purple"},
				{Name: "Terraform", Color: "indigo"},
				{Name: "Ansible", Color: "blue"},
			},
		},
		{
			Title:       "Infrastructure as Code",
			Description: "Established Infrastructure as Code (IaC) practices using Terraform, Ansible, and GitLab CI/CD for VM provisioning, Kubernetes configuration, Route53, and F5 management.",
			Technologies: []models.Technology{
				{Name: "Terraform", Color: "indigo"},
				{Name: "Ansible", Color: "blue"},
				{Name: "GitLab CI/CD", Color: "green"},
			},
		},
		{
			Title:       "VM and Endpoint Life Cycle Management",
			Description: "Managed comprehensive lifecycle operations for VM infrastructure and 12,000+ workstations, including vulnerability remediation and third-party application management, which removed thousands of known exploitable vulnerabilities from the environment.",
			Technologies: []models.Technology{
				{Name: "Datto", Color: "indigo"},
				{Name: "VMware", Color: "red"},
			},
		},
	}

	// Side projects data
	sideProjects := []models.Project{
		{
			Title:       "Dwanium",
			Description: "Service to import Dell PC warranty end dates into Tanium Asset",
			Technologies: []models.Technology{
				{Name: "Rust", Color: "blue"},
				{Name: "Docker", Color: "green"},
				{Name: "Postgres", Color: "purple"},
			},
			Link: "https://github.com/tedewaard/dwanium",
		},
		{
			Title:       "Lenovium",
			Description: "Service to import Lenovo PC warranty end dates into Tanium Asset",
			Technologies: []models.Technology{
				{Name: "Typescript", Color: "yellow"},
			},
			Link: "https://github.com/tedewaard/lenovium",
		},
		{
			Title:       "Edewaard Equipment",
			Description: "Website for Edewaard Equipment - Hudsonville, MI",
			Technologies: []models.Technology{
				{Name: "Typescript", Color: "yellow"},
				{Name: "HTML/CSS", Color: "indigo"},
			},
			Link:        "https://github.com/tedewaard/EER",
			WebsiteLink: "https://edewaardequipment.com",
		},
	}

	// Timeline data
	timeline := []models.TimelineItem{
		{
			Title:       "BI Intern - Haworth",
			Date:        "April 2015",
			Description: "Started my IT career as a Business Intelligence intern at Haworth (Office furniture manufacturer). Worked with SQL databases, created reports and dashboards, and learned the basics of data analysis and visualization.",
			Position:    "right",
		},
		{
			Title:       "Next Internship",
			Date:        "April 2016",
			Description: "BI was interesting but I realized I was lacking an understanding of the fundamentals of IT. Decided to take a new internship at Padnos Recycling Solutions working on the IT Help Desk. Gained hands-on experience with troubleshooting, user support, network infrastructure, and learned the importance of customer service in IT.",
			Position:    "left",
		},
		{
			Title:       "Full Time IT Career",
			Date:        "May 2017",
			Description: "Started working on the IT Service Desk at Service Express. Provided technical support for enterprise customers, managed ticketing systems, and developed strong problem-solving skills in a fast-paced environment.",
			Position:    "right",
		},
		{
			Title:       "Promoted to Systems Administrator",
			Date:        "September 2020",
			Description: "Advanced to Systems Administrator role, taking on more infrastructure responsibilities. Managed server infrastructure, implemented automation solutions, worked with virtualization technologies, and gained experience with cloud platforms.",
			Position:    "left",
		},
		{
			Title:       "IT Analyst at MillerKnoll",
			Date:        "December 2023",
			Description: "Joined MillerKnoll as an IT Analyst to broaden my experience. Worked on enterprise IT projects, collaborated with cross-functional teams, and gained exposure to large-scale corporate IT operations.",
			Position:    "right",
		},
		{
			Title:       "Back at Service Express",
			Date:        "December 2024",
			Description: "Started as an Infrastructure Automation Engineer at Service Express. Focusing on infrastructure as code, CI/CD pipelines, automation frameworks, and modern DevOps practices to streamline operations.",
			Position:    "left",
		},
	}

	// Render the home page with all data
	component := templates.HomePage(workProjects, sideProjects, timeline)
	component.Render(r.Context(), w)
}
