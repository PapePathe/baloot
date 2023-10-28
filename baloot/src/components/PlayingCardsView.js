import { SimpleGrid } from "@chakra-ui/react";
import CardView from "./CardView";

function PlayingCardsView({
  cards,
  onDragStart,
  onDragEnter,
  onDragEnd,
  onClick,
}) {
  return (
    <SimpleGrid
      spacing={2}
      templateColumns="repeat(auto-fill, minmax(120px, 1fr))"
    >
      {cards.map((c) => {
        return c ? (
          <CardView
            onClick={onClick}
            Genre={c.Genre}
            Couleur={c.Couleur}
            onDragStart={(e) => onDragStart(e)}
            onDragEnter={(e) => onDragEnter(e)}
            onDragEnd={onDragEnd}
          />
        ) : null;
      })}
    </SimpleGrid>
  );
}

export default PlayingCardsView;
