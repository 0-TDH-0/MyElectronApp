// Electron main script

const path = require("path");
const { app, BrowserWindow } = require("electron");

function createMainWindow() {
// Create a new browser window
    const window = new BrowserWindow({
    width: 800,
    height: 600
    });

    window.loadFile(path.join(__dirname, "./renderer/index.html"))
}

// Start the Electron application
app.whenReady().then(() => {
    createMainWindow();
});

app.on("window-all-closed", () => {
app.quit();
});