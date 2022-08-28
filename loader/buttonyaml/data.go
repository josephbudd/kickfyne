package buttonyaml

import "github.com/josephbudd/kickfyne/loader"

type YAML struct {
	InsertBeforeButton int           `yaml:"insertBeforeButton"`
	Button             loader.Button `yaml:"button"`
}

const (
	AddButtonYAMLExamples = `HOW TO USE A BUTTON YAML FILE:
◊ A button YAML file let's you add a single button to your application.
◊ You can add a single button after running "$ kickfyne framework".
◊ There are only 2 kinds of buttons.
  1. A button with panels.
  2. A button with tabs.

📌 1. ADD A BUTTON WITH PANELS.

In a button YAML file:
insertBeforeButton:
  ◊ Indicates the zero based index of the button that you want this new button to preceed.
  ◊ In the example YAML file below, the "Personal" button will be inserted before the current first button.
    But since no buttons have been added yet, the "Personal" button will be the first button.
button:
  ◊ Indicates the new button.
  ◊ The button has:
    ⬫ A name.
    ⬫ A label.
    ⬫ A list of panels.
panels:
  ◊ Indicates a list of one or more panels that form a panel group.
  ◊ The first panel listed is the panel group's default panel.
  ◊ Each panel group is completely under your control.
  ◊ In the example YAML file below:
    ⬫ See the Personal button which has 2 panels, Personal and Settings.
  ◊ A panel has:
    ⬫ A name.
    ⬫ A heading which appears at the top of the panel.
    ⬫ A description of the panel's unique purpose.
  ◊ A panel group does not just have your panels. It also has:
    ⬫ A messenger which communicates with the back end using messages.


📄 BUTTON YAML FILE EXAMPLE 1 - A BUTTON WITH PANELS:

insertBeforeButton: 0
button:
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


📌 2. ADD A BUTTON WITH TABS.

As stated above, in a button YAML file:
insertBeforeButton:
  ◊ Indicates the zero based index of the button that you want this new button to preceed.
  ◊ In the example YAML file below, the "Contacts" button will be inserted before the current first button.
    So this "Contacts" button will be inserted before the "Personal" button added in the above example.
    The 2 buttons will be displayed in the following sequence, "Conacts", "Personal".
button:
  ◊ Indicates the new button.
  ◊ The button has:
    ⬫ A name.
    ⬫ A label.
    ⬫ A list of tabs.
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
  ◊ A panel has:
    ⬫ A name.
    ⬫ A heading which appears at the top of the panel.
    ⬫ A description of the panel's unique purpose.
  ◊ A panel group does not just have your panels. It also has:
    ⬫ A messenger which communicates with the back end using messages.

📄 BUTTON YAML FILE EXAMPLE 2 - ADD A BUTTON WITH TABS:
  
insertBeforeButton: 0
button:
  label: Contacts
  name: Contacts
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
`
)
