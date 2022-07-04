game serve based on Leaf server framework
===========
A game server based on [Leaf framework](https://github.com/name5566/leaf).

```
import React, { useEffect, useState } from "react";
import "./App.css";
import { number } from "decoders";
import { Select } from "antd";
const { Option } = Select;
function App() {
  const [data, setData] = useState({});
  const [ws, setWebsocket] = useState<WebSocket>();
  const [isGameStarted, setIsGameStated] = useState(false);
  useEffect(() => {
    const ws = new WebSocket("ws://127.0.0.1:3654");
    setWebsocket(ws);
    ws.binaryType = "arraybuffer";
    ws.onopen = (event) => {
      // Send Hello message
      ws.send(
        JSON.stringify({
          Hello: {
            Name: "Sinak",
          },
        })
      );
    };

    ws.onmessage = function (e) {
      console.log("event:", e?.data?.size);
      var decoder = new window.TextDecoder("utf-8");
      var result = JSON.parse(decoder.decode(e.data));
      console.log(result);
      setData(result);
    };
  }, []);
  useEffect(()=>{
    ws?.send(
      JSON.stringify({
        GameState: {
          status: isGameStarted? 2:1,
        },
      })
    );
  },[isGameStarted])

  return (
    <div className="App">
      <h3>Get started with web socket</h3>
      <h4>Ready to guess!!</h4>
      <pre>{JSON.stringify(data)}</pre>
     
      <br />
      <button
        onClick={() => {
          setIsGameStated(!isGameStarted);
          ws?.send(
            JSON.stringify({
              GameState: {
                status: 1,
              },
            })
          );
        }}
      >
        {!isGameStarted ? "Start Game" : "Stop Game"}
      </button>
      {isGameStarted && (
        <>
        <button
          style={{ marginLeft: "10px" }}
          onClick={() => {
            ws?.send(
              JSON.stringify({
                GuessGame: {
                  text: "big",
                },
              })
            );
         
          }}
        >
          Big
        </button>
        <button
          style={{ marginLeft: "10px" }}
          onClick={() => {
            ws?.send(
              JSON.stringify({
                GuessGame: {
                  text: "small",
                },
              })
            );
         
          }}
        >
          Small
        </button>
        </>
      )}
    </div>
  );
}

export default App;

```

Licensing
---------

Leaf server is licensed under the Apache License, Version 2.0. See [LICENSE](https://github.com/name5566/leafserver/blob/master/LICENSE) for the full license text.
