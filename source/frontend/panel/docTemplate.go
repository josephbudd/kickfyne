package panel

const (
	docFileName = "Doc.go"
)

var docTemplate = `// Package panel is the panel groups.
// PANEL GROUP.
// A panel group is one or more panels.
// Only 1 panel is the group is visible at a time.
// The panel switching in the group is often done by the group's messenger which communicates with the back end.
//
// PANEL GROUP: GROUP.GO FILE.
// The group.go file does 2 things.
// 1. Initializes the group of panels in func Init.
// 2. Returns the group's content in func Content.
//
// THE HOME GROUP.
// The home panel group is the opening panel group.
// It's panel group is at panel/home/group.go.
// The home panel group has only 1 panel and it's content is competely defined in group.go.
// Home only presents buttons.
//  * Each button corresponds to it's own panel group.
//  * When the button is clicked:
//    1. the home panel is hidden.
//    2. it's corresponding panel group is displayed, layed out in a widget/backpanel.
//
// Home.go must
// 1. import each buttons corresponding panel group package.
// 2. In func Init, call the corresponding panel package's Init func.
//   Example:
//     // Setup button.
//     if err = setup.Init(ctx, ctxCancel, app, w); err != nil {
//     	return
//     }
//     // Play button.
//     if err = play.Init(ctx, ctxCancel, app, w); err != nil {
//     	return
//     }
// 3. In func Content
//    func Content() (homeContent *fyne.Container, err error) {
//    
//    	// The onclick for each backpanel's back button.
//    	// backOnClick will cause the window to display this homeContent.
//    	backOnClick := func() {
//    		window.SetContent(homeContent)
//    	}
//    
//    	// The Setup button content.
//    	// Get the content of the setup panel group.
//    	setupGroupContent := setup.Content()
//    	// Make a backpanel using the group content.
//    	// The backpanel's back button will return the person here by displaying homeContent.
//    	setupBackpanelContent := backpanel.Content("Setup", backOnClick, setupGroupContent)
//    	// Create the setup button so that clicking on it sets the window content to the setupBackpanelContent.
//    	setupButton := safebutton.New("Setup", func() { window.SetContent(setupBackpanelContent) })
//    
//    	// The Play button content.
//    	// Get the content of the play panel group.
//    	playGroupContent := play.Content()
//    	// Make a backpanel using the group content.
//    	// The backpanel's back button will return the person here by displaying homeContent.
//    	playBackpanelContent := backpanel.Content("Play", backOnClick, playGroupContent)
//    	// Create the play button so that clicking on it sets the window content to the playBackpanelContent.
//    	playButton := safebutton.New("Play", func() { window.SetContent(playBackpanelContent) })
//    
//    	// Home content.
//    	homeContent = container.NewCenter(container.NewHBox(setupButton, layout.NewSpacer(), playButton))
//    
//    	return
//    }

// 
package panel

`
