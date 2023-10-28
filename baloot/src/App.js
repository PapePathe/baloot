import React, { useState, useCallback, useEffect, useRef } from "react";
import useWebSocket from "react-use-websocket";
import DeckView from "./components/DeckView";
import PlayerCardsView from "./components/PlayerCardsView";
import reorderCards from "./utils/reorderCards";
import messageStore from "./utils/messageStore";
import playCard from "./utils/playCard";
import { Grid, GridItem, VStack, Box } from "@chakra-ui/react";

const WS_URL = "ws://127.0.0.1:7777/ws/100";

function App() {
  const [messageHistory, setMessageHistory] = useState([]);
  const [playingCards, setPlayingCards] = useState([]);
  const [deck, setDeck] = useState([]);
  const [cards, setCards] = useState([]);
  const [takes, setTakes] = useState([]);
  const [gametake, setGametake] = useState(null);
  const [playerTakes, setPlayerTakes] = useState([]);
  const [playerID, setPlayerID] = useState(null);
  const { sendMessage, lastJsonMessage } = useWebSocket(WS_URL);
  const dragItem = useRef();
  const dragOverItem = useRef();
  const dragStart = (e) => {
    dragItem.current = e.target;
  };
  const dragEnter = (e) => {
    dragOverItem.current = e.currentTarget;
  };
  const drop = useCallback(() => {
    reorderCards(dragItem, dragOverItem, cards, setCards);
  }, [dragItem, dragOverItem, cards, setCards]);
  const orderPlayingCards = useCallback(() => {
    reorderCards(dragItem, dragOverItem, playingCards, setPlayingCards);
  }, [dragItem, dragOverItem, playingCards, setPlayingCards]);
  const handleClickSendMessage = useCallback(
    (take, pid) => {
      sendMessage(
        JSON.stringify({ player_id: `${pid}`, gametake: take, id: "2" }),
      );
    },
    [sendMessage],
  );
  const handleClickPlayMessage = useCallback(
    (couleur, genre, event) => {
      playCard(couleur, genre, playerID, sendMessage);
    },
    [playerID, sendMessage],
  );

  useEffect(() => {
    messageStore(
      lastJsonMessage,
      setMessageHistory,
      setDeck,
      setPlayingCards,
      setGametake,
      setPlayerID,
      setPlayerTakes,
      setCards,
      setTakes,
    );
  }, [lastJsonMessage, setMessageHistory]);

  return (
    <div>
      <div>
        <Grid
          h="100vh"
          templateRows="repeat(3, 1fr)"
          templateColumns="repeat(6, 1fr)"
          gap={0}
        >
          <GridItem colSpan={1} rowSpan={3} bg="tomato" />
          <GridItem colSpan={4} rowSpan={3} bg="papayawhip">
            <VStack spacing={0} align="stretch" height="100%">
              <Box h="25%" bg="yellow.200"></Box>
              <DeckView deck={deck} gametake={gametake} />
              <PlayerCardsView
                takes={takes}
                playingCards={playingCards}
                orderPlayingCards={orderPlayingCards}
                cards={cards}
                handleClickSendMessage={handleClickSendMessage}
                handleClickPlayMessage={handleClickPlayMessage}
                dragStart={dragStart}
                dragEnter={dragEnter}
                drop={drop}
                playerID={playerID}
              />
            </VStack>
          </GridItem>
          <GridItem colSpan={1} rowSpan={3} bg="tomato" />
        </Grid>
      </div>
    </div>
  );
}

export default App;
