package panelyaml

import "github.com/josephbudd/kickfyne/loader"

type YAML struct {
	ButtonName string       `yaml:"buttonName"`
	TabName    string       `yaml:"tabName"`
	Panel      loader.Panel `yaml:"panel"`
}

const (
	AddPanelYAMLExample = `HOW TO USE A PANEL YAML FILE:
◊ A panel YAML file let's you add a new panel to a button or tab in your application.
◊ You can only add a panel to;
  1. A button.
  2. A tab.

📌 1. ADD A PANEL TO A BUTTON.

In a panel YAML file:
buttonName:
  ◊ Indicates the name of the button that this panel will be added to.
panel:
  ◊ Indicates a panel.
  ◊ A panel has:
    ⬫ A name.
    ⬫ A heading which appears at the top of the panel.
    ⬫ A description of the panel's unique purpose.

📄 PANEL YAML FILE EXAMPLE 1 - ADD A PANEL TO A BUTTON:

buttonName: Courses
panel:
- name: NotReady
	heading: You don't have any courses to review.
	description: |
	A static text message informing the user that there are no courses to review.
	The messenger will show this page when it receives a CoursesReset message indicating that there are no courses.
	The messenger will show this page when it receives a CoursesPageOf message indicating that there are no courses.


📌 2. ADD A PANEL TO A TAB.

In a panel YAML file:
buttonName:
  ◊ Indicates the name of the button with the tab that this panel will be added to.
tabName:
  ◊ Indicates the name of the tab that is panel is added to.
panel:
  ◊ Indicates a panel.
  ◊ A panel has:
    ⬫ A name.
    ⬫ A heading which appears at the top of the panel.
    ⬫ A description of the panel's unique purpose.
		
📄 PANEL YAML FILE EXAMPLE 2 - ADD A PANEL TO A TAB:

buttonName: Contacts
tabName: Edit
panel:
- name: NotReady
  heading: You don't have any contacts to edit.
  description: |
    A static text message informing the user that there are no contacts to edit.
    The messenger will show this page when it receives a ContactsReset message indicating that there are no contacts.
    The messenger will show this page when it receives a ContactsPageOf message indicating that there are no contacts.
`
)
