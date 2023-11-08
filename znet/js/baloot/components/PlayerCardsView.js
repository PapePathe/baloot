import React from "react";
import TakesGroupView from "./TakesGroupView";
import PlayingCardsView from "./PlayingCardsView";
import { Box } from "@chakra-ui/react";

const PlayerCardsView = ({
  takes,
  playingCards,
  cards,
  handleClickSendMessage,
  playerID,
  orderPlayingCards,
  dragEnter,
  dragStart,
  handleClickPlayMessage,
  drop,
}) => {
  return (
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
          id='playingCards'
        />
      ) : null}
      {cards ? (
        <PlayingCardsView
          cards={cards}
          onDragEnd={drop}
          onDragEnter={dragEnter}
          onDragStart={dragStart}
          id='takingCards'
        />
      ) : null}
    </Box>
  );
};

export default PlayerCardsView;
