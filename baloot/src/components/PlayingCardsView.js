import {
  SimpleGrid
} from '@chakra-ui/react'
import CardView from './CardView';

function PlayingCardsView({cards, onDragStart, onDragEnter, onDragEnd}) {
  return (  <SimpleGrid spacing={0} templateColumns='repeat(auto-fill, minmax(200px, 1fr))'>
    {
      cards.map((c) => {
        return c? (
          <CardView
            Genre={c.Genre}
            Couleur={c.Couleur}
            onDragStart={(e) => onDragStart(e)}
            onDragEnter={(e) => onDragEnter(e)}
            onDragEnd={onDragEnd}
          />
        ) : null
      })
    }
  </SimpleGrid>
  )
}

export default PlayingCardsView
