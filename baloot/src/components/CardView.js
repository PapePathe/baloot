import {SimpleGrid, Button, Card, Heading, CardHeader, CardBody, CardFooter, Header, Text } from '@chakra-ui/react'
import { TrefleIcon, CarreauIcon, PiqueIcon, CoeurIcon } from './Icons';

const icons = {
 "Carreau": <CarreauIcon boxSize={64} color='red.500' />,
  "Coeur": <CoeurIcon boxSize={64} color='red.500' />,
 "Pique": <PiqueIcon boxSize={64} color='red.500' />,
 "Trefle": <TrefleIcon boxSize={64} color='red.500' />,
}

function CardView({ Genre, Couleur, onDragStart, onDragEnter, onDragEnd}) {
  const icon = icons[Couleur];

  return (
    <Card size={'sm'} draggable onDragStart={onDragStart} onDragEnter={onDragEnter} onDragEnd={onDragEnd}>
      <CardHeader>
        <Text textAlign={'left'} fontSize="32" flex='1'>{Genre}</Text>
      </CardHeader>
      <CardBody align='center'>{icon}</CardBody>
      <CardFooter>
        <Text textAlign={'right'} fontSize="32" flex='1'>{Genre}</Text>
      </CardFooter>
    </Card>
  );
}

export default CardView;
