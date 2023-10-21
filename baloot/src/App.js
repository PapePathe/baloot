import React, { useState, useCallback, useEffect, useRef } from 'react';
import useWebSocket, { ReadyState } from 'react-use-websocket';
import {
  Box,
  Button,
  ButtonGroup,
  Card,
  CardHeader,
  CardBody,
  CardFooter,
  Header,
  Heading,
  SimpleGrid,
  Text
} from '@chakra-ui/react'
import CardView from './components/CardView';

const WS_URL = 'ws://127.0.0.1:7777/ws/100';

function App() {
  const [messageHistory, setMessageHistory] = useState([]);
  const { sendMessage, lastMessage, lastJsonMessage, readyState } = useWebSocket(WS_URL);
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

  useEffect(() => {
    if (lastJsonMessage !== null) {
      if (lastJsonMessage !== {}) {
        switch (lastJsonMessage.id) {
          case 1:
            setPlayerID((prev) => lastJsonMessage.player.id)
            setCards((prev) => lastJsonMessage.player.hand.Cards)
            setTakes((prev) => lastJsonMessage.available_takes)
            break
          case 5:
            setPlayerTakes((prev) => [...prev, lastJsonMessage.take])
            setTakes((prev) => lastJsonMessage.available_takes)
            break
        }
      }
      setMessageHistory((prev) => prev.concat(lastMessage));
    }
  }, [lastMessage, setMessageHistory]);

  const handleClickSendMessage = useCallback((take, pid) => {
    console.log(take);
    sendMessage(JSON.stringify({player_id: `${pid}`, gametake: take, id: "2"}))
  }, []);

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
      {takes? (
        <Box m='2'>
          <ButtonGroup size='sm' isAttached variant='outline' spacing={1}>
            {takes.map((t) => {
              return (
                <Button onClick={ (e) => handleClickSendMessage(t.Name, playerID) }>{t.Name}</Button>
              )
            })}
          </ButtonGroup>
        </Box>
      ) : <p>no takes to display</p>}

      {cards? (
        <SimpleGrid spacing={4} templateColumns='repeat(auto-fill, minmax(200px, 1fr))'>
          {cards.map((c) => {
            return (
              <CardView
                Genre={c.Genre}
                Couleur={c.Couleur}
                onDragStart={(e) => dragStart(e)}
                onDragEnter={(e) => dragEnter(e)}
                onDragEnd={drop}
              />
            )
          })}
        </SimpleGrid>
      ) : <p>no cards to display</p>}
   </div>
  );
}



export default App;
