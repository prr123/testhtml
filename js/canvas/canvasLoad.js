function loadCanvas(id) {
        var canvas = document.createElement('canvas');
        div = document.getElementById(id); 
        canvas.id     = "CursorLayer";
        canvas.width  = 1224;
        canvas.height = 768;
        canvas.style.zIndex   = 8;
        canvas.style.position = "absolute";
        canvas.style.border   = "1px solid";
        div.appendChild(canvas)
    }
	
function loadCanvas() {
        var canvas = document.createElement('canvas');
        canvas.id     = "canvGameStage";
        canvas.width  = 1000;
        canvas.height = 768;
        canvas.style.zIndex   = 8;
        canvas.style.position = "absolute";
        canvas.style.border   = "1px solid";
        var div = document.createElement("div");
        div.className = "divGameStage";
        div.appendChild(canvas);
        document.body.appendChild(div)
    }	