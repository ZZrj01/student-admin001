/*
 * @Author: WUMIANHUA
 * @Date: 2021-11-01 08:24:57
 * @LastEditTime: 2021-12-16 21:45:55
 * @Description: 
 */
const {
    app,
    BrowserWindow
} = require('electron')

const path = require('path')
const spawn = require('child_process').spawn

//spawn 不衍生新的线程执行cmd.exe，执行的线程返回的pid就是新的线程pid

//exec 衍生新的线程执行cmd.exe 执行的线程的pid是新的cmd 的pid 不是真实的pid
let loadWindow;

function createWindow() {
    const win = new BrowserWindow({
        minWidth: 950,
        height: 600,
        show: false,
        // electron跨域
        webPreferences: {
            // nodeIntegration: true, //是否集成node
            webSecurity: false
        },
    })
    //开启调试工具
    // win.webContents.openDevTools()
    win.loadFile('index.html')
    win.on('ready-to-show', function () {
        setTimeout(() => {
            loadWindow.close()
            // loadWindow.hide()
            win.show()
        }, 3000)
    })
}

function loadWin() {
    const win = new BrowserWindow({
        width: 400,
        height: 300,
        frame: false
    })
    //开启调试工具
    // win.webContents.openDevTools()
    win.loadFile('load.html')
    //加载完
    loadWindow = win

}

app.whenReady().then(() => {
    runCmd()
    loadWin()
    createWindow()
})

// electron跨域
app.commandLine.appendSwitch('disable-features', 'OutOfBlinkCors');

app.on('window-all-closed', () => {

    //关闭第三方线程
    if (workerProcess != undefined && pid != -1) {
        process.kill(pid)
    }
    app.quit()

})

function runCmd() {
    // 执行
    //cmdPath 服务的名字， pathstr 路劲
    cmdPath = "main.exe"
    let pathStr = path.join(__dirname, "../../server/cmd")
    workerProcess = spawn(cmdPath, [], {
        cwd: pathStr
    })
    // 打印正常的后台可执行程序输出
    workerProcess.stdout.on('data', function (data) {
        console.log("pid:" + workerProcess.pid);
        console.log('stdout: ' + data)
        pid = workerProcess.pid
    })
    // 打印错误的后台可执行程序输出
    workerProcess.stderr.on('data', function (data) {
        console.log('stderr: ' + data)
        pid = -1
    })
    // 退出之后的输出
    workerProcess.on('close', function (code) {
        console.log('out code：' + code)
    })

}