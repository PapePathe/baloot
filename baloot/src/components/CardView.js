import {
  SimpleGrid,
  Button,
  Card,
  Heading,
  CardHeader,
  CardBody,
  CardFooter,
  Header,
  Text,
} from "@chakra-ui/react";
import { TrefleIcon, CarreauIcon, PiqueIcon, CoeurIcon } from "./Icons";

const icons = {
  Carreau: <CarreauIcon boxSize={64} color="red.500" />,
  Coeur: <CoeurIcon boxSize={64} color="red.500" />,
  Pique: <PiqueIcon boxSize={64} color="red.500" />,
  Trefle: <TrefleIcon boxSize={64} color="red.500" />,
};

function CardView({
  Genre,
  Couleur,
  onDragStart,
  onDragEnter,
  onDragEnd,
  onClick,
}) {
  return (
    <Card
      w={128}
      h={64}
      draggable
      onDragStart={onDragStart}
      onDragEnter={onDragEnter}
      onDragEnd={onDragEnd}
      onClick={onClick ? onClick.bind(this, Couleur, Genre) : null}
    >
      <CardHeader>
        <Text textAlign={"left"} fontSize="22" flex="1">
          {Genre}
        </Text>
      </CardHeader>
      <CardBody align="center">{icons[Couleur]}</CardBody>
      <CardFooter>
        <Text textAlign={"right"} fontSize="22" flex="1">
          {Genre}
        </Text>
      </CardFooter>
    </Card>
  );
}

export default CardView;
