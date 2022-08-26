package tabyaml

import "github.com/josephbudd/kickfyne/loader"

type YAML struct {
	ButtonName      string     `yaml:"buttonName"`
	InsertBeforeTab int        `yaml:"insertBeforeTab"`
	Tab             loader.Tab `yaml:"tab"`
}

const (
	AddTabYAMLExample = `HOW TO USE A TAB YAML FILE:
◊ A tab YAML file let's you add a new tab with it's panels, to a button in your application.
◊ The button must have been created or added with tabs not panels.

📌 ADD A TAB.

In a tab YAML file:
buttonName:
  ◊ Indicates the name of the button that this tab will be added to.
insertBeforeTab:
  ◊ Indicates the zero based index of the tab that you want this new tab to preceed.
tab:
  ◊ Indicates the new tab.
  ◊ In the example YAML file below, we are adding the tab named Edit.
  ◊ The Edit tab will be inserted before the current tab at index 2.
  ◊ The tab has:
    ⬫ A name.
    ⬫ A label.
    ⬫ A list of panels.
panels:
  ◊ Indicates a list of one or more panels that form a panel group.
  ◊ The first panel listed is the panel group's default panel.
  ◊ Each panel group is completely under your control.
  ◊ In the example YAML file below:
  ◊ The Edit tab has 3 panels. NotReady, Select and Edit.
  ◊ A panel has:
    ⬫ A name.
    ⬫ A heading which appears at the top of the panel.
    ⬫ A description of the panel's unique purpose.
  ◊ A panel group does not just have your panels. It also has:
    ⬫ A messenger which communicates with the back end using messages.
    ⬫ A stater which communicates with the state.

📄 TAB YAML FILE EXAMPLE:

buttonName: Courses
insertBeforeTab: 2
tab:
- name: Edit
  label: Edit
  panels:
  - name: NotReady
    heading: You don't have any contacts to edit.
    description: |
      A static text message informing the user that there are no contacts to edit.
      The messenger will show this page when it receives a ContactsReset message indicating that there are no contacts.
      The messenger will show this page when it receives a ContactsPageOf message indicating that there are no contacts.
  - name: Select
    heading: Select a contact to edit.
    description: |
      A contact select widget which displays as much info as possible with each contact and allows the user to select a contact.
      The select button will select the contact using the messenger.
      The messenger will show this page when it receives a ContactsReset message indicating that there are contacts.
      The messenger will show this page when it receives a ContactsPageOf message indicating that there are contacts.
      The messenger will initialized this list when it receives the ContactsReset message.
      The messenger will handle page requests from the select widget using the ContactsPageOf message.
      The messenger will request the contact for editing using a ContactForEditing message.
  - name: Edit
    heading: Edit a contact.
    description: |
      A contact form which displays and allows the user to edit a contact.
      The edit button will submit the edits using the messenger.
      The cancel button will go back to the select panel.
      The messenger will fill this form and show this panel upon receiving a ContactForEditing message.
      The messenger will submit the edit using an ContactEdit message.
`
)
