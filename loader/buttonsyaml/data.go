package buttonsyaml

import "github.com/josephbudd/kickfyne/loader"

type YAML struct {
	Buttons []loader.Button `yaml:"buttons"`
}

const (
	CreateButtonsYAMLExample = `HOW TO USE A BUTTONS YAML FILE:
◊ A buttons YAML file will create (not add) multiple buttons, tabs, panels in the application.
◊ The only 3 times you can create buttons are
  1. After building the framework using the command "$ kickfyne framework".
  2. After removing all of the buttons, tab-bars and panel-groups using the command "$ kickfyne frontend clear".
  3. After individually removing each button with it's tab-bar and panel-groups using the command "$ kickfyne frontend button remove <button-name>".
◊ All buttons and tabs and panels remain in the order they appear in the YAML file.
◊ There are only 2 kinds of buttons.
  1. A button with panels.
  2. A button with tabs.

📌 CREATE BUTTONS.

In a buttons YAML file:
buttons:
  ◊ Indicates a list of one or more buttons.
  ◊ In the example YAML file below, there are two buttons; Contacts and Personal.
  ◊ A button has:
    ⬫ A name.
    ⬫ A label.
    ⬫ A list of tabs or panels.
tabs:
  ◊ Indicates a list of one or more tabs that form a tab bar.
  ◊ In the example YAML file below, the Contacts button has 3 tabs; Add, Edit and Remove.
  ◊ A tab has:
    ⬫ A name.
    ⬫ A label.
    ⬫ A list of one or more panels.
panels:
  ◊ Indicates a list of one or more panels that form a panel group.
  ◊ The first panel listed is the panel group's default panel.
  ◊ Each panel group is completely under your control.
  ◊ In the example YAML file below:
    ⬫ See the Contacts button and it's Edit tab which has 3 panels, NotReady, Select and Edit.
    ⬫ See the Personal button which has 2 panels, Personal and Settings.
  ◊ A panel has:
    ⬫ A name.
    ⬫ A heading which appears at the top of the panel.
    ⬫ A description of the panel's unique purpose.
  ◊ A panel group does not just have your panels. It also has:
    ⬫ A messenger which communicates with the back end using messages.
    ⬫ A stater which communicates with the state.

📄 BUTTONS YAML FILE EXAMPLE:

buttons:
- name: Contacts
  label: Contacts
  tabs:
  - name: Add
    label: Add
    panels:
    - name: Add
      heading: Add a new contact.
      description: |
        A contact form.
        The add button will add the contact using the messenger.
        The cancel button will clear the form.	
        The messenger will send the submission to the back end using the ContactAdd message.
        The messenger will clear this page when it receives the ContactAdd message.
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
  - name: Remove
    label: Remove
    panels:
    - name: NotReady
      heading: You don't have any contacts to remove.
      description: |
        A static text message informing the user that there are no contacts to edit.
        The messenger will show this page when it receives a ContactsReset message indicating that there are no contacts.
        The messenger will show this page when it receives a ContactsPageOf message indicating that there are no contacts.
    - name: Select
      heading: Select a contact to remove.
      description: |
        A contact select widget which displays as much info as possible with each contact and allows the user to select a contact.
        The select button will select the contact using the messenger.
        The messenger will show this page when it receives a ContactsReset message indicating that there are contacts.
        The messenger will show this page when it receives a ContactsPageOf message indicating that there are contacts.
        The messenger will initialized this list when it receives a ContactsReset message.
        The messenger will handle page requests from the select widget using a ContactsPageOf message.
        The messenger will select the contact for removal using a ContactForRemoval message.
    - name: Remove
      heading: Remove a contact.
      description: |
        A contact read-only form which displays and allows the user to remove a contact.
        The remove button will submit the contact for removal using the messenger.
        The cancel button will go back to the select panel.
        The messenger will fill this form and show this panel upon receiving a ContactForRemoval message.
        The messenger will submit the removal using a ContactRemove message.
- name: Personal
  label: Personal
  panels:
  - name: Personal
    heading: My personal info.	
    description: |
      A form which displays and allows the user to edit personal information like name, phone, email.
      The edit button will submit the edits using the messenger.
      The cancel button will reset the form.
      The messenger will fill this form at startup using the PersonalForEdit message.
      The messenger will submit the edit using an EditPersonal message.
  - name: Settings
    heading: My personal settings.
    description: |
      A form which displays and allows the user to edit personal settings.
      The edit button will submit the edits using the messenger.
      The cancel button will reset the form.
      The messenger will fill this form at startup using the SettingsForEdit message.
      The messenger will submit the edit using an EditSettings message.
`
)
