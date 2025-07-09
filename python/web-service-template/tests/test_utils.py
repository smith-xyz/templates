"""
Unit tests for the utils module
"""

import pytest
import os
import sys
from unittest.mock import patch

# Add the project root to the Python path
sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..'))

from src.utils import (
    get_project_root,
    validate_environment,
    format_message,
    safe_get_env
)


class TestGetProjectRoot:
    """Test class for get_project_root function."""
    
    def test_get_project_root_returns_string(self):
        """Test that get_project_root returns a string."""
        result = get_project_root()
        assert isinstance(result, str)
        
    def test_get_project_root_is_absolute_path(self):
        """Test that get_project_root returns an absolute path."""
        result = get_project_root()
        assert os.path.isabs(result)


class TestValidateEnvironment:
    """Test class for validate_environment function."""
    
    def test_validate_environment_returns_bool(self):
        """Test that validate_environment returns a boolean."""
        result = validate_environment()
        assert isinstance(result, bool)
    
    @patch('sys.version_info', (3, 8, 0))
    def test_validate_environment_python_38_success(self):
        """Test that Python 3.8 passes validation."""
        result = validate_environment()
        assert result is True
    
    @patch('sys.version_info', (3, 9, 0))
    def test_validate_environment_python_39_success(self):
        """Test that Python 3.9 passes validation."""
        result = validate_environment()
        assert result is True


class TestFormatMessage:
    """Test class for format_message function."""
    
    def test_format_message_without_prefix(self):
        """Test format_message without prefix."""
        message = "Hello, World!"
        result = format_message(message)
        assert result == message
    
    def test_format_message_with_prefix(self):
        """Test format_message with prefix."""
        message = "Hello, World!"
        prefix = "INFO"
        result = format_message(message, prefix)
        assert result == "[INFO] Hello, World!"
    
    def test_format_message_empty_message(self):
        """Test format_message with empty message."""
        result = format_message("", "TEST")
        assert result == "[TEST] "


class TestSafeGetEnv:
    """Test class for safe_get_env function."""
    
    def test_safe_get_env_existing_variable(self):
        """Test safe_get_env with existing environment variable."""
        with patch.dict(os.environ, {'TEST_VAR': 'test_value'}):
            result = safe_get_env('TEST_VAR')
            assert result == 'test_value'
    
    def test_safe_get_env_nonexistent_variable(self):
        """Test safe_get_env with nonexistent environment variable."""
        result = safe_get_env('NONEXISTENT_VAR')
        assert result is None
    
    def test_safe_get_env_with_default(self):
        """Test safe_get_env with default value."""
        result = safe_get_env('NONEXISTENT_VAR', 'default_value')
        assert result == 'default_value'


if __name__ == "__main__":
    pytest.main([__file__]) 