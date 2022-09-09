package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Project struct {
	XMLName                xml.Name               `xml:"project"`
	ModelVersion           string                 `xml:"modelVersion"`
	Parent                 Parent                 `xml:"parent"`
	GroupID                string                 `xml:"groupId"`
	ArtifactID             string                 `xml:"artifactId"`
	Version                string                 `xml:"version"`
	Packaging              string                 `xml:"packaging"`
	Name                   string                 `xml:"name"`
	Description            string                 `xml:"description"`
	URL                    string                 `xml:"url"`
	InceptionYear          string                 `xml:"inceptionYear"`
	Organization           Organization           `xml:"organization"`
	Licenses               []License              `xml:"licenses>license"`
	Developers             []Developer            `xml:"developers>developer"`
	Contributors           []Contributor          `xml:"contributors>contributor"`
	MailingLists           []MailingList          `xml:"mailingLists>mailingList"`
	Prerequisites          Prerequisites          `xml:"prerequisites"`
	Modules                []string               `xml:"modules>module"`
	SCM                    Scm                    `xml:"scm"`
	IssueManagement        IssueManagement        `xml:"issueManagement"`
	CIManagement           CIManagement           `xml:"ciManagement"`
	DistributionManagement DistributionManagement `xml:"distributionManagement"`
	DependencyManagement   DependencyManagement   `xml:"dependencyManagement"`
	Dependencies           []Dependency           `xml:"dependencies>dependency"`
	Repositories           []Repository           `xml:"repositories>repository"`
	PluginRepositories     []PluginRepository     `xml:"pluginRepositories>pluginRepository"`
	Build                  Build                  `xml:"build"`
	Reporting              Reporting              `xml:"reporting"`
	Profiles               []Profile              `xml:"profiles>profile"`
	Properties             Properties             `xml:"properties"`
}

type Properties struct {
	Entries map[string]string
}

func (p *Properties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	type entry struct {
		XMLName xml.Name
		Key     string `xml:"name,attr"`
		Value   string `xml:",chardata"`
	}
	e := entry{}
	p.Entries = map[string]string{}
	for err = d.Decode(&e); err == nil; err = d.Decode(&e) {
		e.Key = e.XMLName.Local
		p.Entries[e.Key] = e.Value
	}
	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

type Parent struct {
	GroupID      string `xml:"groupId"`
	ArtifactID   string `xml:"artifactId"`
	Version      string `xml:"version"`
	RelativePath string `xml:"relativePath"`
}

type Organization struct {
	Name string `xml:"name"`
	URL  string `xml:"url"`
}

type License struct {
	Name         string `xml:"name"`
	URL          string `xml:"url"`
	Distribution string `xml:"distribution"`
	Comments     string `xml:"comments"`
}

type Developer struct {
	ID              string     `xml:"id"`
	Name            string     `xml:"name"`
	Email           string     `xml:"email"`
	URL             string     `xml:"url"`
	Organization    string     `xml:"organization"`
	OrganizationURL string     `xml:"organizationUrl"`
	Roles           []string   `xml:"roles>role"`
	Timezone        string     `xml:"timezone"`
	Properties      Properties `xml:"properties"`
}

type Contributor struct {
	Name            string     `xml:"name"`
	Email           string     `xml:"email"`
	URL             string     `xml:"url"`
	Organization    string     `xml:"organization"`
	OrganizationURL string     `xml:"organizationUrl"`
	Roles           []string   `xml:"roles>role"`
	Timezone        string     `xml:"timezone"`
	Properties      Properties `xml:"properties"`
}

type MailingList struct {
	Name          string   `xml:"name"`
	Subscribe     string   `xml:"subscribe"`
	Unsubscribe   string   `xml:"unsubscribe"`
	Post          string   `xml:"post"`
	Archive       string   `xml:"archive"`
	OtherArchives []string `xml:"otherArchives>otherArchive"`
}

type Prerequisites struct {
	Maven string `xml:"maven"`
}

type Scm struct {
	Connection          string `xml:"connection"`
	DeveloperConnection string `xml:"developerConnection"`
	Tag                 string `xml:"tag"`
	URL                 string `xml:"url"`
}

type IssueManagement struct {
	System string `xml:"system"`
	URL    string `xml:"url"`
}

type CIManagement struct {
	System    string     `xml:"system"`
	URL       string     `xml:"url"`
	Notifiers []Notifier `xml:"notifiers>notifier"`
}

type Notifier struct {
	Type          string     `xml:"type"`
	SendOnError   bool       `xml:"sendOnError"`
	SendOnFailure bool       `xml:"sendOnFailure"`
	SendOnSuccess bool       `xml:"sendOnSuccess"`
	SendOnWarning bool       `xml:"sendOnWarning"`
	Address       string     `xml:"address"`
	Configuration Properties `xml:"configuration"`
}

type DistributionManagement struct {
	Repository         Repository `xml:"repository"`
	SnapshotRepository Repository `xml:"snapshotRepository"`
	Site               Site       `xml:"site"`
	DownloadURL        string     `xml:"downloadUrl"`
	Relocation         Relocation `xml:"relocation"`
	Status             string     `xml:"status"`
}

type Site struct {
	ID   string `xml:"id"`
	Name string `xml:"name"`
	URL  string `xml:"url"`
}

type Relocation struct {
	GroupID    string `xml:"groupId"`
	ArtifactID string `xml:"artifactId"`
	Version    string `xml:"version"`
	Message    string `xml:"message"`
}

type DependencyManagement struct {
	Dependencies []Dependency `xml:"dependencies>dependency"`
}

type Dependency struct {
	GroupID    string      `xml:"groupId"`
	ArtifactID string      `xml:"artifactId"`
	Version    string      `xml:"version"`
	Type       string      `xml:"type"`
	Classifier string      `xml:"classifier"`
	Scope      string      `xml:"scope"`
	SystemPath string      `xml:"systemPath"`
	Exclusions []Exclusion `xml:"exclusions>exclusion"`
	Optional   string      `xml:"optional"`
}

type Exclusion struct {
	ArtifactID string `xml:"artifactId"`
	GroupID    string `xml:"groupId"`
}

type Repository struct {
	UniqueVersion bool             `xml:"uniqueVersion"`
	Releases      RepositoryPolicy `xml:"releases"`
	Snapshots     RepositoryPolicy `xml:"snapshots"`
	ID            string           `xml:"id"`
	Name          string           `xml:"name"`
	URL           string           `xml:"url"`
	Layout        string           `xml:"layout"`
}

type RepositoryPolicy struct {
	Enabled        string `xml:"enabled"`
	UpdatePolicy   string `xml:"updatePolicy"`
	ChecksumPolicy string `xml:"checksumPolicy"`
}

type PluginRepository struct {
	Releases  RepositoryPolicy `xml:"releases"`
	Snapshots RepositoryPolicy `xml:"snapshots"`
	ID        string           `xml:"id"`
	Name      string           `xml:"name"`
	URL       string           `xml:"url"`
	Layout    string           `xml:"layout"`
}

type BuildBase struct {
	DefaultGoal      string           `xml:"defaultGoal"`
	Resources        []Resource       `xml:"resources>resource"`
	TestResources    []Resource       `xml:"testResources>testResource"`
	Directory        string           `xml:"directory"`
	FinalName        string           `xml:"finalName"`
	Filters          []string         `xml:"filters>filter"`
	PluginManagement PluginManagement `xml:"pluginManagement"`
	Plugins          []Plugin         `xml:"plugins>plugin"`
}

type Build struct {
	SourceDirectory       string      `xml:"sourceDirectory"`
	ScriptSourceDirectory string      `xml:"scriptSourceDirectory"`
	TestSourceDirectory   string      `xml:"testSourceDirectory"`
	OutputDirectory       string      `xml:"outputDirectory"`
	TestOutputDirectory   string      `xml:"testOutputDirectory"`
	Extensions            []Extension `xml:"extensions>extension"`
	BuildBase
}

type Extension struct {
	GroupID    string `xml:"groupId"`
	ArtifactID string `xml:"artifactId"`
	Version    string `xml:"version"`
}

type Resource struct {
	TargetPath string   `xml:"targetPath"`
	Filtering  string   `xml:"filtering"`
	Directory  string   `xml:"directory"`
	Includes   []string `xml:"includes>include"`
	Excludes   []string `xml:"excludes>exclude"`
}

type PluginManagement struct {
	Plugins []Plugin `xml:"plugins>plugin"`
}

type Plugin struct {
	GroupID       string            `xml:"groupId"`
	ArtifactID    string            `xml:"artifactId"`
	Version       string            `xml:"version"`
	Extensions    string            `xml:"extensions"`
	Executions    []PluginExecution `xml:"executions>execution"`
	Dependencies  []Dependency      `xml:"dependencies>dependency"`
	Inherited     string            `xml:"inherited"`
	Configuration Configuration     `xml:"configuration"`
}

type Configuration struct {
	Source string `xml:"source"`
	Target string `xml:"target"`
}

type PluginExecution struct {
	ID        string   `xml:"id"`
	Phase     string   `xml:"phase"`
	Goals     []string `xml:"goals>goal"`
	Inherited string   `xml:"inherited"`
}

type Reporting struct {
	ExcludeDefaults string            `xml:"excludeDefaults"`
	OutputDirectory string            `xml:"outputDirectory"`
	Plugins         []ReportingPlugin `xml:"plugins>plugin"`
}

type ReportingPlugin struct {
	GroupID    string      `xml:"groupId"`
	ArtifactID string      `xml:"artifactId"`
	Version    string      `xml:"version"`
	Inherited  string      `xml:"inherited"`
	ReportSets []ReportSet `xml:"reportSets>reportSet"`
}

type ReportSet struct {
	ID        string   `xml:"id"`
	Reports   []string `xml:"reports>report"`
	Inherited string   `xml:"inherited"`
}

type Profile struct {
	ID                     string                 `xml:"id"`
	Activation             Activation             `xml:"activation"`
	Build                  BuildBase              `xml:"build"`
	Modules                []string               `xml:"modules>module"`
	DistributionManagement DistributionManagement `xml:"distributionManagement"`
	Properties             Properties             `xml:"properties"`
	DependencyManagement   DependencyManagement   `xml:"dependencyManagement"`
	Dependencies           []Dependency           `xml:"dependencies>dependency"`
	Repositories           []Repository           `xml:"repositories>repository"`
	PluginRepositories     []PluginRepository     `xml:"pluginRepositories>pluginRepository"`
	Reporting              Reporting              `xml:"reporting"`
}

type Activation struct {
	ActiveByDefault bool               `xml:"activeByDefault"`
	JDK             string             `xml:"jdk"`
	OS              ActivationOS       `xml:"os"`
	Property        ActivationProperty `xml:"property"`
	File            ActivationFile     `xml:"file"`
}

type ActivationOS struct {
	Name    string `xml:"name"`
	Family  string `xml:"family"`
	Arch    string `xml:"arch"`
	Version string `xml:"version"`
}

type ActivationProperty struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

type ActivationFile struct {
	Missing string `xml:"missing"`
	Exists  string `xml:"exists"`
}

func ParsePom(file string) (Project, error) {
	fmt.Println("pom.xml parser..")

	// Open our xmlFile
	xmlFile, err := os.Open("pom.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var project Project
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &project)

	// for i := 0; i < len(project.Build.BuildBase.Plugins); i++ {
	// 	fmt.Println("ArtifactID: " + project.Build.BuildBase.Plugins[i].ArtifactID)
	// 	fmt.Println("Source: " + project.Build.BuildBase.Plugins[i].Configuration.Source)
	// 	fmt.Println("Target: " + project.Build.BuildBase.Plugins[i].Configuration.Target)
	// 	// fmt.Println("User Name: " + users.Users[i].Name)
	// 	// fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	// }
	return project, err
}
