import { Box, Button, ButtonGroup } from "@chakra-ui/react";

function TakesGroupView({ onClickHandler, playerID, takes }) {
  return (
    <Box m="2">
      <ButtonGroup size="sm" isAttached variant="outline" spacing={1}>
        {takes.map((t) => {
          return (
            <Button onClick={(e) => onClickHandler(t.Name, playerID)}>
              {t.Name}
            </Button>
          );
        })}
      </ButtonGroup>
    </Box>
  );
}

export default TakesGroupView;
