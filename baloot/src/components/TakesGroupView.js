import { Box, Button, ButtonGroup } from "@chakra-ui/react";

function TakesGroupView({ onClickHandler, playerID, takes }) {
  return (
    <Box m="2">
      <ButtonGroup size="sm" isAttached variant="outline" spacing={1}>
        {takes.map((t) => {
          return (
            <Button name={t.name} onClick={(e) => onClickHandler(t.name, playerID)}>
              {t.name}
            </Button>
          );
        })}
      </ButtonGroup>
    </Box>
  );
}

export default TakesGroupView;
