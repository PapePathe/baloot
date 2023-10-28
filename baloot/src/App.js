import React, { useState, useCallback, useEffect, useRef } from "react";
import useWebSocket, { ReadyState } from "react-use-websocket";
import TakesGroupView from "./components/TakesGroupView";
import PlayingCardsView from "./components/PlayingCardsView";
import {
  Flex,
  Spacer,
  Text,
  Grid,
  GridItem,
  VStack,
  Box,
  StackDivider,
} from "@chakra-ui/react";

const WS_URL = "ws://127.0.0.1:7777/ws/100";

function App() {
  const [messageHistory, setMessageHistory] = useState([]);
  const { sendMessage, lastJsonMessage, readyState } = useWebSocket(WS_URL);
  const [playingCards, setPlayingCards] = useState([]);
  const [deck, setDeck] = useState([]);
  const [cards, setCards] = useState([]);
  const [takes, setTakes] = useState([]);
  const [gametake, setGametake] = useState(null);
  const [playerTakes, setPlayerTakes] = useState([]);
  const [playerID, setPlayerID] = useState(null);
  const dragItem = useRef();
  const dragOverItem = useRef();

  const dragStart = (e) => {
    dragItem.current = e.target;
  };

  const dragEnter = (e) => {
    dragOverItem.current = e.currentTarget;
  };

  const drop = () => {
    const index = Array.from(dragItem.current.parentNode.children).indexOf(
      dragItem.current,
    );
    const index2 = Array.from(dragOverItem.current.parentNode.children).indexOf(
      dragOverItem.current,
    );
    const copyListItems = [...cards];
    const dragItemCount = copyListItems[index];
    copyListItems.splice(index, 1);
    copyListItems.splice(index2, 0, dragItemCount);
    dragItem.current = null;
    dragOverItem.current = null;
    setCards(copyListItems);
  };

  const orderPlayingCards = () => {
    const index = Array.from(dragItem.current.parentNode.children).indexOf(
      dragItem.current,
    );
    const index2 = Array.from(dragOverItem.current.parentNode.children).indexOf(
      dragOverItem.current,
    );
    const copyListItems = [...playingCards];
    const dragItemCount = copyListItems[index];
    copyListItems.splice(index, 1);
    copyListItems.splice(index2, 0, dragItemCount);
    dragItem.current = null;
    dragOverItem.current = null;
    setPlayingCards(copyListItems);
  };

  useEffect(() => {
    if (lastJsonMessage !== null) {
      if (lastJsonMessage !== {}) {
        switch (lastJsonMessage.id) {
          case 1:
            setPlayerID((prev) => lastJsonMessage.player.id);
            setCards((prev) => lastJsonMessage.player.hand.Cards);
            setTakes((prev) => lastJsonMessage.available_takes);
            break;
          case 2:
            setCards((prev) => []);
            setTakes((prev) => []);
            setPlayerTakes((prev) => []);
            setPlayingCards(
              (prev) => lastJsonMessage.player.playing_hand.Cards,
            );
            setGametake((prev) => lastJsonMessage.gametake.Name);
            break;
          case 5:
            setPlayerTakes((prev) => [...prev, lastJsonMessage.take]);
            setTakes((prev) => lastJsonMessage.available_takes);
            break;
          case 6:
            setDeck((prev) => lastJsonMessage.deck);
            setPlayingCards(
              (prev) => lastJsonMessage.player.playing_hand.Cards,
            );
            break;
          default:
            throw new Error("Error message id not found");
        }
      }
      setMessageHistory((prev) => prev.concat(lastJsonMessage));
    }
  }, [lastJsonMessage, setMessageHistory]);

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
      if (couleur !== "" && genre !== "") {
        sendMessage(
          JSON.stringify({
            player_id: `${playerID}`,
            color: couleur,
            genre: genre,
            id: "4",
          }),
        );
      }
    },
    [playerID, sendMessage],
  );

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
              <Box h="40%" bg="tomato">
                <Flex>
                  <Box p="4">Team a score</Box>
                  <Spacer />
                  <Box p="4">{gametake ? <Text>{gametake}</Text> : null}</Box>
                  <Spacer />
                  <Box p="4">Team b score</Box>
                </Flex>
                {deck ? <PlayingCardsView cards={deck} /> : null}
              </Box>
              <Box h="35%" bg="pink.100">
                {takes ? (
                  <TakesGroupView
                    takes={takes}
                    onClickHandler={handleClickSendMessage}
                    playerID={playerID}
                  />
                ) : null}
                {playingCards ? (
                  <PlayingCardsView
                    cards={playingCards}
                    onDragEnd={orderPlayingCards}
                    onDragEnter={dragEnter}
                    onDragStart={dragStart}
                    onClick={handleClickPlayMessage}
                  />
                ) : null}
                {cards ? (
                  <PlayingCardsView
                    cards={cards}
                    onDragEnd={drop}
                    onDragEnter={dragEnter}
                    onDragStart={dragStart}
                  />
                ) : null}
              </Box>
            </VStack>
          </GridItem>
          <GridItem colSpan={1} rowSpan={3} bg="tomato" />
        </Grid>
      </div>
    </div>
  );
}

export default App;
