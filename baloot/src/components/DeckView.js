import PlayingCardsView from "./PlayingCardsView";
import { Box, Flex, Spacer, Text } from "@chakra-ui/react";

const DeckView = ({ deck, gametake }) => {
  return (
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
  );
};

export default DeckView;
