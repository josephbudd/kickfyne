# kickfyne "fyne my way"

[The Fyne Toolkit](https://fyne.io/) is a favorite application tool kit of mine. After creating a couple of apps I decided to make my own little tool to build and manage building my fyne apps for me. So my tool is called kickfyne.

kickfyne only does 4 things

1. Builds all of my application frameworke code.
1. Builds and manages the GUI a little or a lot at a time.
1. Builds and manages messages and message handlers.
1. Builds and manages data records and data stores.

## Framework summary

All of the application code is contained in 3 folders.

1. frontend/
1. backend/
1. shared/

### The GUI

The GUI begins with a button pad where the user get's a general idea of what the app does based on the buttons.

#### Buttons

There are 2 kinds of buttons.

1. A button that clicks to a tab-bar.
1. A button that clicks to a group of panels.

#### Tabs

Each tab in a tab-bar also has a group of panels.

#### Panels and panel groups

* Only 1 panel is a panel group is visible at a time.
* You control each panel you create.
* Each group of panels also has a messenger which communicates with the backend using by sending and receving messages.
* Each group of panels also has a state which communicates with the state. The state is something that the backend can read from and write to but the front-end can only read from.

So the GUI is built by adding buttons, tabs and panels to the front end.

kickfyne lets me:

## Build the application code

```shell
$ kickfyne framework
```

## Build the front end one step at a time

### Create multiple buttons with tab-bars and or panel-groups

```shell
$ kickfyne frontend create <path to create-buttons-tabs-panels-yaml-file>
```

or

### Add a single button with a group of panels

```shell
$ kickfyne frontend button add <path to add-button-panels-yaml-file>
```

or

### Add a single button with a tab-bar and each tab's group of panels

```shell
$ kickfyne frontend button add <path to add-button-tabs-panels-yaml-file>
```

or

### Add a tab and it's group of panels to a button's tab-bar

```shell
$ kickfyne frontend tab add <path to add-tab-yaml-file>
```

or

### Add a panel to a button's group of panels

```shell
$ kickfyne frontend panel add <path to add-panel-yaml-file>
```

or

### Remove a button along with it's tab-bars and groups of panels

```shell
$ kickfyne frontend button remove <button-name>
```

or

### Remove a tab along with it's group of panels.

```shell
$ kickfyne frontend tab remove <button-name> <tab-name>
```

or

### Remove a panel from a button

```shell
$ kickfyne frontend panel remove <button-name> <panel-name>
```

or

### Remove a panel from a tab

```shell
$ kickfyne frontend panel remove <button-name> <tab-name> <panel-name>
```

## Build messages so that the front-end and back-end can communicate

When a message is added, the back-end message handler is also created in the backend/txrx/ folder.

```shell
$ kickfyne message add <message-name>
```

```shell
$ kickfyne message remove <message-name>
```

```shell
$ kickfyne message list
```

## Build records and stores for storing and retriving data

Adding a record creates:

* The record definition in shared/store/record/.
* The store interface definition in shared/store/storer/.
* The store interface implementation in shared/store/storing/.

```shell
$ kickfyne record add <record-name>
```

```shell
$ kickfyne record remove <record-name>
```

```shell
$ kickfyne record list
```

## Inline Instructions

All GO files generated by kickfyne contain instructions at **KICKYNE TODO** comments.