const initializeController = function (chip8) {
    const {
        getPixel,
        onKeyPress,
        loadRom,
        getClockRate,
        increaseClockRate,
        decreaseClockRate,
        shouldDraw,
        setEmulatorState,
        nextInstruction,
        getScreen,
    } = chip8;

    function updateDisplay() {
        const pixels = JSON.parse(getScreen());
        if (shouldDraw()) {
            for (let i = 0; i < 32; i++) {
                for (let j = 0; j < 64; j++) {
                    const pixel = pixels[i][j];
                    const key = `pixel-${i}-${j}`;
                    const pixelUI = document.getElementById(key);
                    if (pixelUI) {
                        pixelUI.className = pixel ? "on" : "off";
                    }
                }
            }
        }
        requestAnimationFrame(updateDisplay);
    }
    function handleMouseDown(evt) {
        const key = Number(evt.target.id.split("-")[1]);
        onKeyPress(key);
    }
    function handleMouseUp() {
        onKeyPress(255);
    }
    function setupClockRateControls() {
        const clockRateControl = document.getElementById("clock-rate-control");
        clockRateControl.addEventListener("click", function (evt) {
            const { id } = evt.target;
            console.log(id);
            if (id === "btn-increase-cr") increaseClockRate();
            if (id === "btn-decrease-cr") decreaseClockRate();
            const clockRateSpan = document.getElementById("clock-rate-display");
            clockRateSpan.innerHTML = getClockRate();
        });
    }
    function setupEmuControls() {
        const emuControl = document.getElementById("emu-control");
        emuControl.addEventListener("click", function (evt) {
            const { id } = evt.target;
            console.log(id);
            if (id === "btn-pause") setEmulatorState("PAUSED");
            if (id === "btn-start") setEmulatorState("RUNNING");
            if (id === "btn-next") {
                const prevInstruction = document.getElementById("prev-instruction");
                const currentInstruction = document.getElementById("current-instruction");
                const nextInstructionUI = document.getElementById("next-instruction");
                const state = nextInstruction();
                const registers = JSON.parse(state.state0.registers);
                prevInstruction.innerText = currentInstruction.innerText;
                currentInstruction.innerText = state.state0.instruction;
                nextInstructionUI.innerText = state.state1.instruction;
                for (let i = 0; i < 15; i++) {
                    const gpReg = document.getElementById(`v${i}`);
                    gpReg.value = registers.GeneralPurpose[i];
                }
                const iReg = document.getElementById(`i-reg`);
                iReg.value = registers.I;
                const pc = document.getElementById(`pc-reg`);
                pc.value = registers.ProgramCounter;
                const stackPtr = document.getElementById(`stack-reg`);
                stackPtr.value = registers.StackPtr;
            }
        });
    }
    function createRegisterInput(container, id, text) {
        const label = document.createElement("label");
        const input = document.createElement("input");
        const textNode = document.createTextNode(text);
        input.id = id;
        input.disabled = true;
        label.for = id;
        label.appendChild(textNode);
        container.appendChild(label);
        container.appendChild(input);
    }
    function setupRegistersControls() {
        const gpRegs = document.getElementById("gp-regs");
        //set vx regs
        for (let i = 0; i < 16; i++) {
            const div = document.createElement("div");
            div.className = "reg";
            const label = document.createElement("label");
            const reg = document.createElement("input");
            const id = `v${i}`;
            const text = document.createTextNode(id);
            reg.id = id;
            reg.disabled = true;
            label.for = id;
            label.appendChild(text);
            div.appendChild(label);
            div.appendChild(reg);
            gpRegs.appendChild(div);
        }
        createRegisterInput(document.getElementById("i-container"), "i-reg", "I");
        createRegisterInput(document.getElementById("pc-container"), "pc-reg", "PC");
        createRegisterInput(document.getElementById("stack-container"), "stack-reg", "Stack Ptr");
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
                loadRom(bytes);
            };
            reader.readAsArrayBuffer(this.files[0]);
            requestAnimationFrame(updateDisplay);
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
            setupEmuControls();
            setupRegistersControls();
        },
    };
};
