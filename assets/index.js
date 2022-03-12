const chip8API = (function () {
    function updateDisplay() {
        for (let i = 0; i < 32; i++) {
            for (let j = 0; j < 64; j++) {
                const pixel = getPixel(j, i);
                const key = `pixel-${i}-${j}`;
                const pixelUI = document.getElementById(key);
                if (pixelUI) {
                    pixelUI.className = pixel ? "on" : "off";
                }
            }
        }
    }
    function handleMouseDown(evt) {
        const key = Number(evt.target.id.split("-")[1]);
        onKeypress(key);
    }
    function handleMouseUp() {
        onKeypress(255);
    }
    function setupClockRateControls() {
        const clockRateControl = document.getElementById("clock-rate-control");
        clockRateControl.addEventListener("click", function (evt) {
            const { id } = evt.target;
            console.log(id);
            if (id === "btn-increase-cr") increaseClockSpeed();
            if (id === "btn-decrease-cr") decreaseClockSpeed();
            const clockRateSpan = document.getElementById("clock-rate-display");
            clockRateSpan.innerHTML = getClockSpeed();
        });
    }
    function setupDisplay() {
        const display = document.getElementById("chip8Display");
        for (let i = 0; i < 32; i++) {
            for (let j = 0; j < 64; j++) {
                let pixelUI = document.createElement("div");
                pixelUI.id = `pixel-${i}-${j}`;
                pixelUI.className = "off";
                display.append(pixelUI);
            }
        }
        document.getElementById("chip8ROMUploader").addEventListener("change", function () {
            const reader = new FileReader();
            reader.onload = (ev) => {
                bytes = new Uint8Array(ev.target.result);
                loadROM(bytes);
            };
            reader.readAsArrayBuffer(this.files[0]);
            setInterval(updateDisplay, 10);
        });
    }
    function setupKeyboard() {
        const keys = [
            { id: "key-31", label: "1" },
            { id: "key-32", label: "2" },
            { id: "key-33", label: "3" },
            { id: "key-43", label: "C" },
            { id: "key-34", label: "4" },
            { id: "key-35", label: "5" },
            { id: "key-36", label: "6" },
            { id: "key-44", label: "D" },
            { id: "key-37", label: "7" },
            { id: "key-38", label: "8" },
            { id: "key-39", label: "9" },
            { id: "key-45", label: "E" },
            { id: "key-41", label: "A" },
            { id: "key-0", label: "0" },
            { id: "key-42", label: "B" },
            { id: "key-43", label: "F" },
        ];
        const chip8Keyboard = document.getElementById("chip8Keyboard");
        chip8Keyboard.addEventListener("mouseup", handleMouseUp);
        chip8Keyboard.addEventListener("mousedown", handleMouseDown);
        for (const key of keys) {
            const button = document.createElement("button");
            button.id = key.id;
            const text = document.createTextNode(key.label);
            button.appendChild(text);
            chip8Keyboard.appendChild(button);
        }
    }
    return {
        init: function () {
            setupDisplay();
            setupKeyboard();
            setupClockRateControls();
        },
    };
})();
