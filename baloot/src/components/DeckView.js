import PlayingCardsView from "./PlayingCardsView";
import { Box, Flex, Spacer, Text } from "@chakra-ui/react";

const DeckView = ({ deck, gametake, score }) => {
  return (
    <Box h="40%" bg="tomato">
      <Flex>
        <Box p="4">{score[0]}</Box>
        <Spacer />
        <Box p="4">{gametake ? <Text>{gametake}</Text> : null}</Box>
        <Spacer />
        <Box p="4">{score[1]}</Box>
      </Flex>
      {deck ? <PlayingCardsView cards={deck} id='deckCards' /> : null}
    </Box>
  );
};

export default DeckView;
