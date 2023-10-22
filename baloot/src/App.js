import React, { useState, useCallback, useEffect, useRef } from 'react';
import useWebSocket, { ReadyState } from 'react-use-websocket';
import TakesGroupView from './components/TakesGroupView';
import PlayingCardsView from './components/PlayingCardsView';

const WS_URL = 'ws://127.0.0.1:7777/ws/100';

function App() {
  const [messageHistory, setMessageHistory] = useState([]);
  const { sendMessage, lastJsonMessage, readyState } = useWebSocket(WS_URL);
  const [ playingCards, setPlayingCards] = useState([]);
  const [ cards, setCards] = useState([]);
  const [ takes, setTakes] = useState([]);
  const [ playerTakes, setPlayerTakes] = useState([]);
  const [ playerID, setPlayerID] = useState(null);
  const dragItem = useRef();
  const dragOverItem = useRef();

  const dragStart = e => {
    dragItem.current = e.target
  }

  const dragEnter = e => {
    dragOverItem.current = e.currentTarget
  }

  const drop = () => {
    const index = Array.from(dragItem.current.parentNode.children).indexOf(dragItem.current);
    const index2 = Array.from(dragOverItem.current.parentNode.children).indexOf(dragOverItem.current);
    const copyListItems = [...cards];
    const dragItemCount = copyListItems[index];
    copyListItems.splice(index, 1);
    copyListItems.splice(index2, 0, dragItemCount);
    dragItem.current = null;
    dragOverItem.current = null;
    setCards(copyListItems)
  }

  const orderPlayingCards = () => {
    const index = Array.from(dragItem.current.parentNode.children).indexOf(dragItem.current);
    const index2 = Array.from(dragOverItem.current.parentNode.children).indexOf(dragOverItem.current);
    const copyListItems = [...playingCards];
    const dragItemCount = copyListItems[index];
    copyListItems.splice(index, 1);
    copyListItems.splice(index2, 0, dragItemCount);
    dragItem.current = null;
    dragOverItem.current = null;
    setPlayingCards(copyListItems)
  }

  useEffect(() => {
    if (lastJsonMessage !== null) {
      if (lastJsonMessage !== {}) {
        console.log(lastJsonMessage);
        switch (lastJsonMessage.id) {
          case 1:
            setPlayerID((prev) => lastJsonMessage.player.id)
            setCards((prev) => lastJsonMessage.player.hand.Cards)
            setTakes((prev) => lastJsonMessage.available_takes)
            break
          case 2:
            setCards((prev) => [])
            setTakes((prev) => [])
            setPlayerTakes((prev) => ['test'])
            setPlayingCards((prev) => lastJsonMessage.player.playing_hand.Cards)
            break
          case 5:
            setPlayerTakes((prev) => [...prev, lastJsonMessage.take])
            setTakes((prev) => lastJsonMessage.available_takes)
            break
          default:
            throw new Error("Error message id not found")
        }
      }
      setMessageHistory((prev) => prev.concat(lastJsonMessage));
    }
  }, [lastJsonMessage, setMessageHistory]);

  const handleClickSendMessage = useCallback((take, pid) => {
    console.log(take);
    sendMessage(JSON.stringify({player_id: `${pid}`, gametake: take, id: "2"}))
  }, [sendMessage]);

  const connectionStatus = {
    [ReadyState.CONNECTING]: 'Connecting',
    [ReadyState.OPEN]: 'Open',
    [ReadyState.CLOSING]: 'Closing',
    [ReadyState.CLOSED]: 'Closed',
    [ReadyState.UNINSTANTIATED]: 'Uninstantiated',
  }[readyState];

  return (
    <div>
      <span>The WebSocket is currently {connectionStatus}</span>
      {takes.length > 0 ? (
        <TakesGroupView takes={takes} onClickHandler={handleClickSendMessage} playerID={playerID}  />
      ) : null }

      {playingCards? (
        <PlayingCardsView cards={playingCards} onDragEnd={orderPlayingCards} onDragEnter={dragEnter} onDragStart={dragStart}  />
      ) : null }

      {cards? (
        <PlayingCardsView cards={cards} onDragEnd={drop} onDragEnter={dragEnter} onDragStart={dragStart}  />
      ) : null}
   </div>
  );
}



export default App;
