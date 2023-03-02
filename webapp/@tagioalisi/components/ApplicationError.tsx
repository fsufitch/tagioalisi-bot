import {
  Card,
  CardContent,
  CardHeader,
  Container,
  Typography,
  CardActions,
  Button,
  Stack,
} from '@mui/material';
import React, { useTransition } from 'react';
import { isRouteErrorResponse, useNavigate, useRouteError } from 'react-router-dom';
import { ErrorResponse } from '@remix-run/router';
import ErrorIcon from '@mui/icons-material/Error';
import LinkOffIcon from '@mui/icons-material/LinkOff';
import { CodeBlock } from './CodeBlock';
import HouseIcon from '@mui/icons-material/House';
import UndoIcon from '@mui/icons-material/Undo';

// See: https://reactrouter.com/en/main/start/tutorial#handling-not-found-errors

export const ApplicationError = () => {
  const error = useRouteError();
  return (
    <Container maxWidth="md">
      <Stack direction="column" justifyContent="center" minHeight="95vh">
        {isRouteErrorResponse(error) ? (
          <RouteError error={error} />
        ) : (
          <GenericError error={error} />
        )}
      </Stack>
    </Container>
  );
};

export default ApplicationError;

const RouteError: React.FC<{
  error: ErrorResponse;
}> = ({ error }) => {
  const navigate = useNavigate();
  const [pending, startTransition] = useTransition();

  return (
    <Card>
      <CardHeader
        avatar={<LinkOffIcon />}
        title={
          <h2>
            Routing error ({error.status}): {error.statusText}
          </h2>
        }
      />
      <CardContent>
        <Typography>You flummoxed the page router!</Typography>
        <CodeBlock>{error.data}</CodeBlock>
      </CardContent>
      <CardActions css={{ justifyContent: 'right' }}>
        <Button
          startIcon={<UndoIcon />}
          disabled={pending}
          onClick={() => startTransition(() => navigate(-1))}
        >
          Go Back
        </Button>
        <Button
          startIcon={<HouseIcon />}
          disabled={pending}
          onClick={() => startTransition(() => navigate('/'))}
        >
          Home
        </Button>
      </CardActions>
    </Card>
  );
};

const GenericError: React.FC<{ error: unknown }> = ({ error }) => {
  const navigate = useNavigate();
  const [pending, startTransition] = useTransition();
  console.error('Application error:', error);
  return (
    <Card>
      <CardHeader avatar={<ErrorIcon />} title={<h2>Fatal application error</h2>} />
      <CardContent>
        <Typography>
          This error is really weird, and almost certainly indicates a bug in the page.
        </Typography>
        <CodeBlock>{`${error}`}</CodeBlock>
      </CardContent>
      <CardActions css={{ justifyContent: 'right' }}>
        <Button
          startIcon={<HouseIcon />}
          disabled={pending}
          onClick={() => startTransition(() => navigate('/'))}
        >
          Home
        </Button>
      </CardActions>
    </Card>
  );
};
